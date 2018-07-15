package sclient

import (
	"encoding/json"
	"errors"
	"strings"
)

// Response Object
type Response struct {
	object interface{}
}

// ParseJSON Return response
func ParseJSON(sample []byte) (*Response, error) {
	var Response = Response{}

	if err := json.Unmarshal(sample, &Response.object); err != nil {
		return nil, err
	}

	return &Response, nil
}

func (g *Response) ChildrenMap() (map[string]*Response, error) {
	if mmap, ok := g.Data().(map[string]interface{}); ok {
		children := map[string]*Response{}
		for name, obj := range mmap {
			children[name] = &Response{obj}
		}
		return children, nil
	}
	return nil, ErrNotObj
}
func (g *Response) Children() ([]*Response, error) {
	if array, ok := g.Data().([]interface{}); ok {
		children := make([]*Response, len(array))
		for i := 0; i < len(array); i++ {
			children[i] = &Response{array[i]}
		}
		return children, nil
	}
	if mmap, ok := g.Data().(map[string]interface{}); ok {
		children := []*Response{}
		for _, obj := range mmap {
			children = append(children, &Response{obj})
		}
		return children, nil
	}
	return nil, ErrNotObjOrArray
}

// Path chilren response
func (g *Response) Path(path string) *Response {
	return g.Search(strings.Split(path, ".")...)
}

// Serach chilren response
func (g *Response) Search(hierarchy ...string) *Response {
	var object interface{}

	object = g.Data()
	for target := 0; target < len(hierarchy); target++ {
		if mmap, ok := object.(map[string]interface{}); ok {
			object, ok = mmap[hierarchy[target]]
			if !ok {
				return nil
			}
		} else if marray, ok := object.([]interface{}); ok {
			tmpArray := []interface{}{}
			for _, val := range marray {
				tmpGabs := &Response{val}
				res := tmpGabs.Search(hierarchy[target:]...)
				if res != nil {
					tmpArray = append(tmpArray, res.Data())
				}
			}
			if len(tmpArray) == 0 {
				return nil
			}
			return &Response{tmpArray}
		} else {
			return nil
		}
	}
	return &Response{object}
}

// Data is return data response
func (g *Response) Data() interface{} {
	if g == nil {
		return nil
	}
	return g.object
}

var (
	// ErrOutOfBounds - Index out of bounds.
	ErrOutOfBounds = errors.New("out of bounds")

	// ErrNotObjOrArray - The target is not an object or array type.
	ErrNotObjOrArray = errors.New("not an object or array")

	// ErrNotObj - The target is not an object type.
	ErrNotObj = errors.New("not an object")

	// ErrNotArray - The target is not an array type.
	ErrNotArray = errors.New("not an array")

	// ErrPathCollision - Creating a path failed because an element collided with an existing value.
	ErrPathCollision = errors.New("encountered value collision whilst building path")

	// ErrInvalidInputObj - The input value was not a map[string]interface{}.
	ErrInvalidInputObj = errors.New("invalid input object")

	// ErrInvalidInputText - The input data could not be parsed.
	ErrInvalidInputText = errors.New("input text could not be parsed")

	// ErrInvalidPath - The filepath was not valid.
	ErrInvalidPath = errors.New("invalid file path")

	// ErrInvalidBuffer - The input buffer contained an invalid JSON string
	ErrInvalidBuffer = errors.New("input buffer contained invalid JSON")
)
