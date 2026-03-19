package output

import (
	"encoding/json"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
	"gopkg.in/yaml.v3"
)

func PrintTickets(tickets []models.Ticket, format string) error {
	switch format {
	case "json":
		return printJSON(tickets)
	case "yaml":
		return printYAML(tickets)
	default:
		return printTicketsTable(tickets)
	}
}

func PrintTicket(ticket *models.Ticket, format string) error {
	switch format {
	case "json":
		return printJSON(ticket)
	case "yaml":
		return printYAML(ticket)
	default:
		return printTicketTable(ticket)
	}
}

func PrintFullTicket(fullTicket *models.FullTicket, format string) error {
	switch format {
	case "json":
		return printJSON(fullTicket)
	case "yaml":
		return printYAML(fullTicket)
	default:
		return printFullTicketTable(fullTicket)
	}
}

func PrintContacts(contacts []models.Contact, format string) error {
	switch format {
	case "json":
		return printJSON(contacts)
	case "yaml":
		return printYAML(contacts)
	default:
		return printContactsTable(contacts)
	}
}

func PrintContact(contact *models.Contact, format string) error {
	switch format {
	case "json":
		return printJSON(contact)
	case "yaml":
		return printYAML(contact)
	default:
		return printContactTable(contact)
	}
}

func PrintComments(comments []models.Comment, format string) error {
	switch format {
	case "json":
		return printJSON(comments)
	case "yaml":
		return printYAML(comments)
	default:
		return printCommentsTable(comments)
	}
}

func PrintComment(comment *models.Comment, format string) error {
	switch format {
	case "json":
		return printJSON(comment)
	case "yaml":
		return printYAML(comment)
	default:
		return printCommentTable(comment)
	}
}

func PrintJSON(v interface{}) error {
	return printJSON(v)
}

func PrintDepartments(departments []models.Department, format string) error {
	switch format {
	case "json":
		return printJSON(departments)
	case "yaml":
		return printYAML(departments)
	default:
		return printDepartmentsTable(departments)
	}
}

func PrintDepartment(department *models.Department, format string) error {
	switch format {
	case "json":
		return printJSON(department)
	case "yaml":
		return printYAML(department)
	default:
		return printDepartmentTable(department)
	}
}

func PrintAgents(agents []models.Agent, format string) error {
	switch format {
	case "json":
		return printJSON(agents)
	case "yaml":
		return printYAML(agents)
	default:
		return printAgentsTable(agents)
	}
}

func PrintAgent(agent *models.Agent, format string) error {
	switch format {
	case "json":
		return printJSON(agent)
	case "yaml":
		return printYAML(agent)
	default:
		return printAgentTable(agent)
	}
}

func PrintTicketContext(context *models.TicketContext, format string) error {
	switch format {
	case "json":
		return printJSON(context)
	case "yaml":
		return printYAML(context)
	default:
		return printTicketContextTable(context)
	}
}

func PrintTimeEntries(entries []models.TimeEntry, format string) error {
	switch format {
	case "json":
		return printJSON(entries)
	case "yaml":
		return printYAML(entries)
	default:
		return printTimeEntriesTable(entries)
	}
}

func PrintTimeEntry(entry *models.TimeEntry, format string) error {
	switch format {
	case "json":
		return printJSON(entry)
	case "yaml":
		return printYAML(entry)
	default:
		return printTimeEntryTable(entry)
	}
}

func PrintTimeReport(report *models.TimeReport, format string) error {
	switch format {
	case "json":
		return printJSON(report)
	case "yaml":
		return printYAML(report)
	default:
		return printTimeReportTable(report)
	}
}

func PrintAttachments(attachments []models.Attachment, format string) error {
	switch format {
	case "json":
		return printJSON(attachments)
	case "yaml":
		return printYAML(attachments)
	default:
		return printAttachmentsTable(attachments)
	}
}

func PrintAttachment(attachment *models.Attachment, format string) error {
	switch format {
	case "json":
		return printJSON(attachment)
	case "yaml":
		return printYAML(attachment)
	default:
		return printAttachmentTable(attachment)
	}
}

func PrintArticles(articles []models.Article, format string) error {
	switch format {
	case "json":
		return printJSON(articles)
	case "yaml":
		return printYAML(articles)
	default:
		return printArticlesTable(articles)
	}
}

func PrintArticle(article *models.Article, format string) error {
	switch format {
	case "json":
		return printJSON(article)
	case "yaml":
		return printYAML(article)
	default:
		return printArticleTable(article)
	}
}

func PrintCategories(categories []models.Category, format string) error {
	switch format {
	case "json":
		return printJSON(categories)
	case "yaml":
		return printYAML(categories)
	default:
		return printCategoriesTable(categories)
	}
}

func PrintCategory(category *models.Category, format string) error {
	switch format {
	case "json":
		return printJSON(category)
	case "yaml":
		return printYAML(category)
	default:
		return printCategoryTable(category)
	}
}

func PrintTasks(tasks []models.Task, format string) error {
	switch format {
	case "json":
		return printJSON(tasks)
	case "yaml":
		return printYAML(tasks)
	default:
		return printTasksTable(tasks)
	}
}

func PrintTask(task *models.Task, format string) error {
	switch format {
	case "json":
		return printJSON(task)
	case "yaml":
		return printYAML(task)
	default:
		return printTaskTable(task)
	}
}

func PrintTicketStats(stats *models.TicketStats, format string) error {
	switch format {
	case "json":
		return printJSON(stats)
	case "yaml":
		return printYAML(stats)
	default:
		return printTicketStatsTable(stats)
	}
}

func PrintAgentStats(agents []models.AgentStats, format string) error {
	switch format {
	case "json":
		return printJSON(agents)
	case "yaml":
		return printYAML(agents)
	default:
		return printAgentStatsTable(agents)
	}
}

func PrintSLAStats(stats *models.SLAStats, format string) error {
	switch format {
	case "json":
		return printJSON(stats)
	case "yaml":
		return printYAML(stats)
	default:
		return printSLAStatsTable(stats)
	}
}

func printJSON(v interface{}) error {
	encoder := json.NewEncoder(output)
	encoder.SetIndent("", "  ")
	return encoder.Encode(v)
}

func printYAML(v interface{}) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Print(string(data))
	return nil
}