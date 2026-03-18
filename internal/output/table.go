package output

import (
	"fmt"
	"os"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
	"github.com/olekukonko/tablewriter"
)

var output = os.Stdout

func printTicketsTable(tickets []models.Ticket) error {
	if len(tickets) == 0 {
		fmt.Println("No tickets found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Number", "Subject", "Status", "Priority", "Created"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, t := range tickets {
		subject := t.Subject
		if len(subject) > 50 {
			subject = subject[:47] + "..."
		}
		table.Append([]string{
			t.ID,
			t.TicketNumber,
			subject,
			t.Status,
			t.Priority,
			t.CreatedTime[:10],
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d tickets\n", len(tickets))
	return nil
}

func printTicketTable(ticket *models.Ticket) error {
	fmt.Printf("\nTicket: %s (%s)\n", ticket.TicketNumber, ticket.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Subject:     %s\n", ticket.Subject)
	fmt.Printf("Status:       %s\n", ticket.Status)
	fmt.Printf("Priority:     %s\n", ticket.Priority)
	fmt.Printf("Created:  %s\n", ticket.CreatedTime)
	fmt.Printf("Modified: %s\n", ticket.ModifiedTime)
	if ticket.ContactEmail != "" {
		fmt.Printf("Contact:      %s <%s>\n", ticket.ContactName, ticket.ContactEmail)
	}
	if ticket.DepartmentName != "" {
		fmt.Printf("Department:   %s\n", ticket.DepartmentName)
	}
	if ticket.AssigneeName != "" {
		fmt.Printf("Assignee:     %s\n", ticket.AssigneeName)
	}
	if ticket.Description != "" {
		fmt.Println("\nDescription:")
		fmt.Println(ticket.Description)
	}
	return nil
}

func printFullTicketTable(fullTicket *models.FullTicket) error {
	printTicketTable(&fullTicket.Ticket)

	if fullTicket.Contact != nil {
		fmt.Println("\nContact:")
		fmt.Printf("  Name:  %s %s\n", fullTicket.Contact.FirstName, fullTicket.Contact.LastName)
		fmt.Printf("  Email: %s\n", fullTicket.Contact.Email)
		if fullTicket.Contact.Phone != "" {
			fmt.Printf("  Phone: %s\n", fullTicket.Contact.Phone)
		}
	}

	if len(fullTicket.Threads) > 0 {
		fmt.Printf("\nThreads (%d):\n", len(fullTicket.Threads))
		for i, thread:= range fullTicket.Threads {
			fmt.Printf("  %d. [%s] %s\n", i+1, thread.CreatedTime[:10], thread.ThreadType)
			if thread.Subject != "" {
				fmt.Printf("     Subject: %s\n", thread.Subject)
			}
			if thread.Description != "" {
				desc := thread.Description
				if len(desc) >100 {
					desc = desc[:97] + "..."
				}
				fmt.Printf("     %s\n", desc)
			}
		}
	}

	return nil
}

func printContactsTable(contacts []models.Contact) error {
	if len(contacts) == 0{
		fmt.Println("No contacts found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name", "Email", "Phone"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, c := range contacts {
		name := fmt.Sprintf("%s %s", c.FirstName, c.LastName)
		table.Append([]string{
			c.ID,
			name,
			c.Email,
			c.Phone,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d contacts\n", len(contacts))
	return nil
}

func printContactTable(contact *models.Contact) error {
	fmt.Printf("\nContact: %s\n", contact.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Name:       %s %s\n", contact.FirstName, contact.LastName)
	fmt.Printf("Email:      %s\n", contact.Email)
	if contact.Phone != "" {
		fmt.Printf("Phone:      %s\n", contact.Phone)
	}
	if contact.Mobile != "" {
		fmt.Printf("Mobile:     %s\n", contact.Mobile)
	}
	if contact.Department != "" {
		fmt.Printf("Department: %s\n", contact.Department)
	}
	if contact.Title != "" {
		fmt.Printf("Title:      %s\n", contact.Title)
	}
	fmt.Printf("Created:    %s\n", contact.CreatedTime)
	return nil
}

func printCommentsTable(comments []models.Comment) error {
	if len(comments) == 0 {
		fmt.Println("No comments found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Author", "Type", "Created", "Preview"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, c := range comments {
		preview := c.Content
		if len(preview) > 40 {
			preview = preview[:37] + "..."
		}
		commentType := "Private"
		if c.IsPublic {
			commentType = "Public"
		}
		table.Append([]string{
			c.ID,
			c.AuthorName,
			commentType,
			c.CreatedTime[:10],
			preview,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d comments\n", len(comments))
	return nil
}

func printCommentTable(comment *models.Comment) error {
	fmt.Printf("\nComment: %s\n", comment.ID)
	fmt.Println("─────────────────────────────────────")
	commentType := "Private"
	if comment.IsPublic {
		commentType = "Public"
	}
	fmt.Printf("Type:       %s\n", commentType)
	fmt.Printf("Author:     %s\n", comment.AuthorName)
	fmt.Printf("Created:    %s\n", comment.CreatedTime)
	fmt.Println("\nContent:")
	fmt.Println(comment.Content)
	return nil
}

func printBatchResultJSON(result *models.BatchResponse) error {
	return printJSON(result)
}