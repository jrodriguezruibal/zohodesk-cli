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