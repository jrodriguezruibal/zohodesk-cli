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

func printDepartmentsTable(departments []models.Department) error {
	if len(departments) == 0 {
		fmt.Println("No departments found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name", "Visible"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, d := range departments {
		visible := "No"
		if d.IsVisible {
			visible = "Yes"
		}
		table.Append([]string{
			d.ID,
			d.Name,
			visible,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d departments\n", len(departments))
	return nil
}

func printDepartmentTable(department *models.Department) error {
	fmt.Printf("\nDepartment: %s\n", department.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Name:        %s\n", department.Name)
	if department.Description != "" {
		fmt.Printf("Description: %s\n", department.Description)
	}
	fmt.Printf("Visible:     %v\n", department.IsVisible)
	return nil
}

func printAgentsTable(agents []models.Agent) error {
	if len(agents) == 0 {
		fmt.Println("No agents found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name", "Email", "Active"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, a := range agents {
		active := "No"
		if a.IsActive {
			active = "Yes"
		}
		table.Append([]string{
			a.ID,
			a.Name,
			a.Email,
			active,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d agents\n", len(agents))
	return nil
}

func printAgentTable(agent *models.Agent) error {
	fmt.Printf("\nAgent: %s\n", agent.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Name:   %s\n", agent.Name)
	fmt.Printf("Email:  %s\n", agent.Email)
	if agent.Role != "" {
		fmt.Printf("Role:   %s\n", agent.Role)
	}
	if agent.Phone != "" {
		fmt.Printf("Phone:  %s\n", agent.Phone)
	}
	active := "Inactive"
	if agent.IsActive {
		active = "Active"
	}
	fmt.Printf("Status: %s\n", active)
	return nil
}

func printTicketContextTable(context *models.TicketContext) error {
	printTicketTable(context.Ticket)

	if context.Contact != nil {
		fmt.Println("\nContact:")
		fmt.Printf("  Name:  %s %s\n", context.Contact.FirstName, context.Contact.LastName)
		fmt.Printf("  Email: %s\n", context.Contact.Email)
	}

	if context.Department != nil {
		fmt.Println("\nDepartment:")
		fmt.Printf("  Name: %s\n", context.Department.Name)
	}

	if context.Assignee != nil {
		fmt.Println("\nAssignee:")
		fmt.Printf("  Name:  %s\n", context.Assignee.Name)
		fmt.Printf("  Email: %s\n", context.Assignee.Email)
	}

	if len(context.Threads) > 0 {
		fmt.Printf("\nThreads (%d):\n", len(context.Threads))
		for i, thread := range context.Threads {
			fmt.Printf("  %d. [%s] %s\n", i+1, thread.CreatedTime[:10], thread.ThreadType)
		}
	}

	return nil
}

func printTimeEntriesTable(entries []models.TimeEntry) error {
	if len(entries) == 0 {
		fmt.Println("No time entries found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Agent", "Duration", "Description", "Executed"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, e := range entries {
		duration := formatDuration(e.Hours, e.Minutes)
		description := e.Description
		if len(description) > 30 {
			description = description[:27] + "..."
		}
		executed := ""
		if e.ExecutedTime != "" {
			executed = e.ExecutedTime[:10]
		} else if e.CreatedTime != "" {
			executed = e.CreatedTime[:10]
		}
		table.Append([]string{
			e.ID,
			e.AgentName,
			duration,
			description,
			executed,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d time entries\n", len(entries))
	return nil
}

func printTimeEntryTable(entry *models.TimeEntry) error {
	fmt.Printf("\nTime Entry: %s\n", entry.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Duration:    %s\n", formatDuration(entry.Hours, entry.Minutes))
	fmt.Printf("Agent:       %s\n", entry.AgentName)
	if entry.Description != "" {
		fmt.Printf("Description: %s\n", entry.Description)
	}
	if entry.ExecutedTime != "" {
		fmt.Printf("Executed:    %s\n", entry.ExecutedTime)
	}
	if entry.TicketID != "" {
		fmt.Printf("Ticket ID:   %s\n", entry.TicketID)
	}
	fmt.Printf("Created:     %s\n", entry.CreatedTime)
	return nil
}

func printTimeReportTable(report *models.TimeReport) error {
	fmt.Println("\nTime Report Summary")
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Total Time:     %.1f hours\n", report.TotalHours)
	fmt.Printf("Total Entries:  %d\n", report.EntriesCount)

	if len(report.EntriesByAgent) > 0 {
		fmt.Println("\nBy Agent:")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{"Agent", "Hours", "Entries"})
		table.SetBorder(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		for _, a := range report.EntriesByAgent {
			table.Append([]string{
				a.AgentName,
				fmt.Sprintf("%.1f", a.Hours),
				fmt.Sprintf("%d", a.Count),
			})
		}
		table.Render()
	}

	if len(report.EntriesByTicket) > 0 {
		fmt.Println("\nBy Ticket:")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{"Ticket", "Subject", "Hours", "Entries"})
		table.SetBorder(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		for _, t := range report.EntriesByTicket {
			subject := t.Subject
			if len(subject) > 30 {
				subject = subject[:27] + "..."
			}
			table.Append([]string{
				t.TicketNumber,
				subject,
				fmt.Sprintf("%.1f", t.Hours),
				fmt.Sprintf("%d", t.Count),
			})
		}
		table.Render()
	}

	return nil
}

func formatDuration(hours, minutes int) string {
	if hours > 0 && minutes > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%dh", hours)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm", minutes)
	}
	return "0m"
}

func printAttachmentsTable(attachments []models.Attachment) error {
	if len(attachments) == 0 {
		fmt.Println("No attachments found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name", "Size", "Type"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, a := range attachments {
		size := formatSize(a.Size)
		table.Append([]string{
			a.ID,
			a.Name,
			size,
			a.ContentType,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d attachments\n", len(attachments))
	return nil
}

func printAttachmentTable(attachment *models.Attachment) error {
	fmt.Printf("\nAttachment: %s\n", attachment.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Name:  %s\n", attachment.Name)
	fmt.Printf("Size:  %s\n", formatSize(attachment.Size))
	fmt.Printf("Type:  %s\n", attachment.ContentType)
	if attachment.DownloadURL != "" {
		fmt.Printf("URL:   %s\n", attachment.DownloadURL)
	}
	return nil
}

func formatSize(bytes int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.1f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.1f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.1f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

func printArticlesTable(articles []models.Article) error {
	if len(articles) == 0 {
		fmt.Println("No articles found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Title", "Status", "Author", "Created"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, a := range articles {
		title := a.Title
		if len(title) > 40 {
			title = title[:37] + "..."
		}
		created := ""
		if a.CreatedTime != "" {
			created = a.CreatedTime[:10]
		}
		table.Append([]string{
			a.ID,
			title,
			a.Status,
			a.AuthorName,
			created,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d articles\n", len(articles))
	return nil
}

func printArticleTable(article *models.Article) error {
	fmt.Printf("\nArticle: %s\n", article.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Title:    %s\n", article.Title)
	if article.Summary != "" {
		fmt.Printf("Summary:  %s\n", article.Summary)
	}
	fmt.Printf("Status:   %s\n", article.Status)
	fmt.Printf("Author:   %s\n", article.AuthorName)
	if article.CategoryName != "" {
		fmt.Printf("Category: %s\n", article.CategoryName)
	}
	fmt.Printf("Views:    %d\n", article.ViewCount)
	fmt.Printf("Likes:    %d\n", article.LikeCount)
	fmt.Printf("Created:  %s\n", article.CreatedTime)
	if article.ModifiedTime != "" {
		fmt.Printf("Modified: %s\n", article.ModifiedTime)
	}
	if len(article.Tags) > 0 {
		fmt.Printf("Tags:     %v\n", article.Tags)
	}
	if article.Content != "" {
		fmt.Println("\nContent:")
		fmt.Println(article.Content)
	}
	return nil
}

func printCategoriesTable(categories []models.Category) error {
	if len(categories) == 0 {
		fmt.Println("No categories found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name", "Articles"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, c := range categories {
		table.Append([]string{
			c.ID,
			c.Name,
			fmt.Sprintf("%d", c.ArticleCount),
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d categories\n", len(categories))
	return nil
}

func printCategoryTable(category *models.Category) error {
	fmt.Printf("\nCategory: %s\n", category.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Name:     %s\n", category.Name)
	if category.Description != "" {
		fmt.Printf("Description: %s\n", category.Description)
	}
	fmt.Printf("Articles: %d\n", category.ArticleCount)
	fmt.Printf("Created:  %s\n", category.CreatedTime)
	return nil
}

func printTasksTable(tasks []models.Task) error {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Title", "Status", "Priority", "Owner", "Due"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, t := range tasks {
		title := t.Title
		if len(title) > 30 {
			title = title[:27] + "..."
		}
		due := ""
		if t.DueDate != "" {
			due = t.DueDate[:10]
		}
		table.Append([]string{
			t.ID,
			title,
			t.Status,
			t.Priority,
			t.OwnerName,
			due,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d tasks\n", len(tasks))
	return nil
}

func printTaskTable(task *models.Task) error {
	fmt.Printf("\nTask: %s\n", task.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Title:    %s\n", task.Title)
	fmt.Printf("Status:   %s\n", task.Status)
	if task.Description != "" {
		fmt.Printf("Description: %s\n", task.Description)
	}
	if task.Priority != "" {
		fmt.Printf("Priority: %s\n", task.Priority)
	}
	if task.OwnerName != "" {
		fmt.Printf("Owner:   %s\n", task.OwnerName)
	}
	if task.DueDate != "" {
		fmt.Printf("Due:     %s\n", task.DueDate)
	}
	if task.TicketID != "" {
		fmt.Printf("Ticket:  %s\n", task.TicketID)
	}
	fmt.Printf("Created: %s\n", task.CreatedTime)
	if task.CompletedAt != "" {
		fmt.Printf("Completed: %s\n", task.CompletedAt)
	}
	return nil
}

func printTicketStatsTable(stats *models.TicketStats) error {
	fmt.Println("\nTicket Statistics")
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Total:       %d\n", stats.Total)
	fmt.Printf("Open:        %d\n", stats.Open)
	fmt.Printf("Closed:      %d\n", stats.Closed)
	fmt.Printf("On Hold:     %d\n", stats.OnHold)
	fmt.Printf("In Progress: %d\n", stats.InProgress)

	if len(stats.ByStatus) > 0 {
		fmt.Println("\nBy Status:")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{"Status", "Count"})
		table.SetBorder(false)
		for _, s := range stats.ByStatus {
			table.Append([]string{s.Status, fmt.Sprintf("%d", s.Count)})
		}
		table.Render()
	}

	if len(stats.ByPriority) > 0 {
		fmt.Println("\nBy Priority:")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{"Priority", "Count"})
		table.SetBorder(false)
		for _, p := range stats.ByPriority {
			table.Append([]string{p.Priority, fmt.Sprintf("%d", p.Count)})
		}
		table.Render()
	}

	return nil
}

func printAgentStatsTable(agents []models.AgentStats) error {
	if len(agents) == 0 {
		fmt.Println("No agent statistics found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"Agent", "Assigned", "Resolved", "Avg Response"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, a := range agents {
		table.Append([]string{
			a.AgentName,
			fmt.Sprintf("%d", a.TicketsAssigned),
			fmt.Sprintf("%d", a.TicketsResolved),
			a.AvgResponseTime,
		})
	}

	table.Render()
	fmt.Printf("\nTotal agents: %d\n", len(agents))
	return nil
}

func printSLAStatsTable(stats *models.SLAStats) error {
	fmt.Println("\nSLA Compliance")
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Total Tickets:   %d\n", stats.TotalTickets)
	fmt.Printf("Complied:        %d\n", stats.Complied)
	fmt.Printf("Violated:        %d\n", stats.Violated)
	fmt.Printf("Compliance Rate: %.1f%%\n", stats.ComplianceRate)

	if len(stats.ByPriority) > 0 {
		fmt.Println("\nBy Priority:")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{"Priority", "Total", "Complied", "Rate"})
		table.SetBorder(false)
		for _, p := range stats.ByPriority {
			table.Append([]string{
				p.Priority,
				fmt.Sprintf("%d", p.Total),
				fmt.Sprintf("%d", p.Complied),
				fmt.Sprintf("%.1f%%", p.Rate),
			})
		}
		table.Render()
	}

	return nil
}

func printProductsTable(products []models.Product) error {
	if len(products) == 0 {
		fmt.Println("No products found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name", "Code", "Active"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, p := range products {
		active := "No"
		if p.IsActive {
			active = "Yes"
		}
		table.Append([]string{
			p.ID,
			p.Name,
			p.ProductCode,
			active,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d products\n", len(products))
	return nil
}

func printProductTable(product *models.Product) error {
	fmt.Printf("\nProduct: %s\n", product.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Name:  %s\n", product.Name)
	if product.ProductCode != "" {
		fmt.Printf("Code:  %s\n", product.ProductCode)
	}
	if product.Description != "" {
		fmt.Printf("Description: %s\n", product.Description)
	}
	if product.Category != "" {
		fmt.Printf("Category: %s\n", product.Category)
	}
	active := "Inactive"
	if product.IsActive {
		active = "Active"
	}
	fmt.Printf("Status: %s\n", active)
	fmt.Printf("Created: %s\n", product.CreatedTime)
	return nil
}

func printAccountsTable(accounts []models.Account) error {
	if len(accounts) == 0 {
		fmt.Println("No accounts found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name", "Email", "Type", "Active"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, a := range accounts {
		active := "No"
		if a.IsActive {
			active = "Yes"
		}
		table.Append([]string{
			a.ID,
			a.Name,
			a.Email,
			a.Type,
			active,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d accounts\n", len(accounts))
	return nil
}

func printAccountTable(account *models.Account) error {
	fmt.Printf("\nAccount: %s\n", account.ID)
	fmt.Println("─────────────────────────────────────")
	fmt.Printf("Name:  %s\n", account.Name)
	if account.Email != "" {
		fmt.Printf("Email: %s\n", account.Email)
	}
	if account.Phone != "" {
		fmt.Printf("Phone: %s\n", account.Phone)
	}
	if account.Website != "" {
		fmt.Printf("Website: %s\n", account.Website)
	}
	if account.Type != "" {
		fmt.Printf("Type:     %s\n", account.Type)
	}
	if account.Industry != "" {
		fmt.Printf("Industry: %s\n", account.Industry)
	}
	if account.OwnerName != "" {
		fmt.Printf("Owner:   %s\n", account.OwnerName)
	}
	active := "Inactive"
	if account.IsActive {
		active = "Active"
	}
	fmt.Printf("Status: %s\n", active)
	fmt.Printf("Created: %s\n", account.CreatedTime)
	return nil
}

func printTagsTable(tags []models.Tag) error {
	if len(tags) == 0 {
		fmt.Println("No tags found")
		return nil
	}

	table := tablewriter.NewWriter(output)
	table.SetHeader([]string{"ID", "Name"})
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, t := range tags {
		table.Append([]string{
			t.ID,
			t.Name,
		})
	}

	table.Render()
	fmt.Printf("\nTotal: %d tags\n", len(tags))
	return nil
}