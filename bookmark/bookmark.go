package bookmark

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Bookmark represents a single bookmark with a URL and title
type Bookmark struct {
	Title    string
	URL      string
	Children []Bookmark // For nested folders
}

type BookmarkFolder struct {
	Name      string
	Bookmarks []Bookmark
	Children  []*BookmarkFolder
}

type BookmarkService struct {
	importMap map[string]struct{}
}

func NewBookmarkService() *BookmarkService {
	return &BookmarkService{
		importMap: make(map[string]struct{}),
	}
}

func (s *BookmarkService) Import(ctx context.Context, file io.Reader) ([]*BookmarkFolder, error) {
	doc, err := s.readHTMLFile(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading HTML file: %v\n", err)
	}

	bookmarks, err := s.parseBookmarks(doc)
	if err != nil {
		return nil, fmt.Errorf("Error parsing bookmarks: %v\n", err)
	}

	if len(bookmarks) == 0 || len(bookmarks[0].Children) == 0 || len(bookmarks[0].Children[0].Children) == 0 {
		return nil, errors.New("no bookmarks found")
	}

	return s.importBookmark(ctx, nil, bookmarks[0].Children[0].Children), nil
}

func (s *BookmarkService) importBookmark(ctx context.Context, folder *BookmarkFolder, bookmarks []Bookmark) []*BookmarkFolder {
	output := make([]*BookmarkFolder, 0)
	for _, bookmark := range bookmarks {
		if bookmark.URL == "" { //create folder

			if _, ok := s.importMap[bookmark.Title]; ok {
				continue
			}

			s.importMap[bookmark.Title] = struct{}{}

			var nextFolder *BookmarkFolder
			if folder == nil {
				nextFolder = &BookmarkFolder{Name: bookmark.Title}
			} else {
				folder.Children = append(folder.Children, &BookmarkFolder{
					Name: bookmark.Title,
				})
				nextFolder = folder.Children[len(folder.Children)-1]
			}

			output = append(output, nextFolder)

			//import bookmark to bookmark's folder
			_ = s.importBookmark(ctx, nextFolder, bookmark.Children)
		} else { //import bookmark

			if _, ok := s.importMap[bookmark.URL]; ok {
				continue
			}

			if folder == nil {
				folder = &BookmarkFolder{Name: ""}
				output = append(output, folder)
			}

			folder.Bookmarks = append(folder.Bookmarks, bookmark)

			s.importMap[bookmark.URL] = struct{}{}
		}
	}

	return output
}

// parseBookmark parses a single bookmark and returns the title and URL
func (s *BookmarkService) parseBookmark(n *html.Node) (string, string) {
	var title, url string

	// Search for <A> tag which contains the bookmark information
	for _, attr := range n.Attr {
		if strings.ToUpper(attr.Key) == "HREF" {
			url = attr.Val
		}
	}

	// Search for the title inside <A> or <DT> text node
	if n.Type == html.TextNode {
		title = strings.TrimSpace(n.Data)
	} else if n.Type == html.ElementNode {
		title = n.Parent.FirstChild.FirstChild.Data
	}

	return title, url
}

// parseBookmarks recursively traverses the HTML node tree to extract bookmarks
func (s *BookmarkService) parseBookmarks(n *html.Node) ([]Bookmark, error) {
	var bookmarks []Bookmark

	if n.Type == html.ElementNode && strings.ToUpper(n.Data) == "DT" {
		// Check for <A> element inside <DT> element
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && strings.ToUpper(c.Data) == "A" {
				title, url := s.parseBookmark(c)
				if title != "" && url != "" {
					bookmarks = append(bookmarks, Bookmark{Title: title, URL: url})
				}
			}
		}
	}

	// Check for <DL> element (which could be a folder)
	if n.Type == html.ElementNode && strings.ToUpper(n.Data) == "DL" {
		// Recurse into child nodes (could be other folders or bookmarks)
		var folderBookmarks []Bookmark
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			childBookmarks, err := s.parseBookmarks(c)
			if err != nil {
				return nil, err
			}
			folderBookmarks = append(folderBookmarks, childBookmarks...)
		}

		// If the folder has any bookmarks, add them as children
		if len(folderBookmarks) > 0 {
			// Create a dummy folder node, we will set the title later
			bookmarks = append(bookmarks, Bookmark{Title: n.Parent.FirstChild.FirstChild.Data, Children: folderBookmarks})
		}
	}

	// Recurse into child nodes (e.g., folders with nested bookmarks)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childBookmarks, err := s.parseBookmarks(c)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, childBookmarks...)
	}

	return bookmarks, nil
}

// readHTMLFile reads the HTML file containing bookmarks and returns the root node
func (s *BookmarkService) readHTMLFile(file io.Reader) (*html.Node, error) {

	// Parse the HTML file into a node tree
	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
