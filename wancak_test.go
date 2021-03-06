package wancak

import (
	"testing"
)

func TestGetSectionPosts(t *testing.T) {
	posts, err := GetPosts("lol", "")
	if err != nil {
		t.Errorf("Error gettting posts: %v", err)
	}
	postLen := len(posts.Posts)
	if postLen != 3 {
		t.Errorf("Expected post length 3, got %d", postLen)
	}
}

func TestInvalidPageId(t *testing.T) {
	_, err := GetPosts("lol", "123123")
	if err != NotFoundErr {
		t.Errorf("Expected not found error, got %v", err)
	}
}

func TestGetSinglePost(t *testing.T) {
	post, err := GetPostId("1009441")
	if err != nil {
		t.Errorf("Shouldn't be an error, but got error: %v", err)
	}
	if post.Title != "Mungkin Rooney Sudah Lelah Main Bola" {
		t.Errorf("Expected \"Mungkin Rooney Sudah Lelah Main Bola\", got %s", post.Title)
	}
	if post.Id != "1009441" {
		t.Error("Expected 1009441, but ", post.Id)
	}
	if post.NSFW {
		t.Error("This post is not nsfw")
	}
	if post.Img != "http://cdn1.1cak.tv/posts/d7c311f7790379833f042e16cb113800_t.jpg" {
		t.Error("Wrong img url, got", post.Img)
	}
	if post.Url != "http://1cak.com/1009441" {
		t.Errorf("Expected http://1cak.com/1009441, got %s", post.Url)
	}
}

func TestNsfwPost(t *testing.T) {
	post, err := GetPostId("1009496")
	if err != nil {
		t.Errorf("Shouldn't be an error, but got error: %v", err)
	}
	if !post.NSFW {
		t.Error("This is nsfw post")
	}
	if post.Img != "http://1cak.com/images/unsave.jpg" {
		t.Errorf("Wrong img url, got %s", post.Img)
	}
}

func TestGetInvalidPostId(t *testing.T) {
	_, err := GetPostId("123")
	if err != NotFoundErr {
		t.Errorf("Expected not found error, got %v", err)
	}
}

func TestSectionValidation(t *testing.T) {
	if _, err := GetPosts("lawl", ""); err != InvalidSectionErr {
		t.Errorf("Expected sectionErr, got %v", err)
	}
}
