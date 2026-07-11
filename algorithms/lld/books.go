package lld

import (
	"fmt"
	"time"
)

type Genre int

const (
	Fiction Genre = iota
	NonFiction
	Science
)

type BookStatus int

const (
	Available BookStatus = iota
	Borrowed
)

type Book struct {
	ID       string
	Name     string
	Genre    Genre
	Borrower *Member
	Status   BookStatus
}

type Member struct {
	ID    string
	Name  string
	Books []*Book
}

type BorrowRecord struct {
	Book    *Book
	Member  *Member
	DueDate time.Time
}

type Library struct {
	Books   []*Book
	Members []*Member
}

func (b *Book) IsAvailable() bool {
	if b.Status == Available {
		return true
	}
	return false
}

func (l *Library) BorrowBook(member *Member, b *Book) *BorrowRecord {
	// check if book is available
	// if yes:
	//   set book status to Borrowed
	//   set book borrower to member
	//   add book to member's Books slice
	//   create and return BorrowRecord with due date 14 days from now
	// if no: return nil
	if b.IsAvailable() {
		b.Status = Borrowed
		b.Borrower = member
		member.Books = append(member.Books, b)
		return &BorrowRecord{
			Book:    b,
			Member:  member,
			DueDate: time.Now().Add(14 * 24 * time.Hour),
		}

	}
	return nil
}

func (l *Library) ReturnBook(record *BorrowRecord) {
	// 1. set book status to Available
	// 2. set book borrower to nil
	// 3. remove book from member's Books slice
	record.Book.Status = Available
	record.Book.Borrower = nil
	for i, b := range record.Member.Books {
		if b == record.Book {
			record.Member.Books = append(record.Member.Books[:i], record.Member.Books[i+1:]...)
			break
		}
	}
}

func Demo() {
	book := &Book{
		ID:     "B001",
		Name:   "Go Programming",
		Genre:  Fiction,
		Status: Available,
	}
	member := &Member{
		ID:   "M001",
		Name: "John",
	}

	library := &Library{
		Books:   []*Book{book},
		Members: []*Member{member},
	}

	record := library.BorrowBook(member, book)
	fmt.Printf("Book borrowed, due: %v\n", record.DueDate)
	library.ReturnBook(record)
	fmt.Println("After return, available:", book.IsAvailable())

}
