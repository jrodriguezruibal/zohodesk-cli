package cmd

import (
	"context"
	"fmt"

	"github.com/jrodriguezruibal/zohodesk-cli/internal/api"
	"github.com/jrodriguezruibal/zohodesk-cli/internal/output"
	"github.com/jrodriguezruibal/zohodesk-cli/pkg/models"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage tasks",
	Long:  `Manage tasks associated with tickets.`,
}

var tasksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  `List all tasks or tasks for a specific ticket.`,
	RunE:  runTasksList,
}

var tasksGetCmd = &cobra.Command{
	Use:   "get <task-id>",
	Short: "Get task details",
	Long:  `Get detailed information about a specific task.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTasksGet,
}

var tasksCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long:  `Create a new task, optionally associated with a ticket.`,
	RunE:  runTasksCreate,
}

var tasksUpdateCmd = &cobra.Command{
	Use:   "update <task-id>",
	Short: "Update a task",
	Long:  `Update task details.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTasksUpdate,
}

var tasksCompleteCmd = &cobra.Command{
	Use:   "complete <task-id>",
	Short: "Mark a task as complete",
	Long:  `Mark a task as completed.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTasksComplete,
}

var tasksDeleteCmd = &cobra.Command{
	Use:   "delete <task-id>",
	Short: "Delete a task",
	Long:  `Delete a task.`,
	Args:  cobra.ExactArgs(1),
	RunE:  runTasksDelete,
}

var (
	flagTaskTicketID  string
	flagTaskTitle     string
	flagTaskDescription string
	flagTaskPriority  string
	flagTaskOwner     string
	flagTaskDueDate   string
	flagTaskStatus    string
)

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.AddCommand(tasksListCmd)
	tasksCmd.AddCommand(tasksGetCmd)
	tasksCmd.AddCommand(tasksCreateCmd)
	tasksCmd.AddCommand(tasksUpdateCmd)
	tasksCmd.AddCommand(tasksCompleteCmd)
	tasksCmd.AddCommand(tasksDeleteCmd)

	tasksListCmd.Flags().StringVarP(&flagTaskTicketID, "ticket", "t", "", "filter by ticket ID")

	tasksCreateCmd.Flags().StringVarP(&flagTaskTitle, "title", "t", "", "task title (required)")
	tasksCreateCmd.Flags().StringVarP(&flagTaskDescription, "description", "d", "", "task description")
	tasksCreateCmd.Flags().StringVarP(&flagTaskPriority, "priority", "P", "", "task priority")
	tasksCreateCmd.Flags().StringVarP(&flagTaskOwner, "owner", "o", "", "owner ID")
	tasksCreateCmd.Flags().StringVarP(&flagTaskDueDate, "due", "D", "", "due date (YYYY-MM-DD)")
	tasksCreateCmd.Flags().StringVarP(&flagTaskTicketID, "ticket", "T", "", "associated ticket ID")

	tasksUpdateCmd.Flags().StringVarP(&flagTaskTitle, "title", "t", "", "task title")
	tasksUpdateCmd.Flags().StringVarP(&flagTaskDescription, "description", "d", "", "task description")
	tasksUpdateCmd.Flags().StringVarP(&flagTaskPriority, "priority", "P", "", "task priority")
	tasksUpdateCmd.Flags().StringVarP(&flagTaskOwner, "owner", "o", "", "owner ID")
	tasksUpdateCmd.Flags().StringVarP(&flagTaskDueDate, "due", "D", "", "due date (YYYY-MM-DD)")
	tasksUpdateCmd.Flags().StringVarP(&flagTaskStatus, "status", "s", "", "task status")
}

func runTasksList(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	tasks, err := api.NewTasksService(client).List(context.Background(), flagTaskTicketID)
	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTasks(tasks, out)
}

func runTasksGet(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	taskID := args[0]

	task, err := api.NewTasksService(client).Get(context.Background(), taskID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTask(task, out)
}

func runTasksCreate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	if flagTaskTitle == "" {
		return fmt.Errorf("--title is required")
	}

	req := models.TaskCreateRequest{
		Title:       flagTaskTitle,
		Description: flagTaskDescription,
		Priority:    flagTaskPriority,
		OwnerID:     flagTaskOwner,
		DueDate:     flagTaskDueDate,
		TicketID:    flagTaskTicketID,
	}

	task, err := api.NewTasksService(client).Create(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTask(task, out)
}

func runTasksUpdate(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	taskID := args[0]

	req := models.TaskUpdateRequest{
		Title:       flagTaskTitle,
		Description: flagTaskDescription,
		Priority:    flagTaskPriority,
		OwnerID:     flagTaskOwner,
		DueDate:     flagTaskDueDate,
		Status:      flagTaskStatus,
	}

	task, err := api.NewTasksService(client).Update(context.Background(), taskID, req)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	out := getOutputFormat()
	return output.PrintTask(task, out)
}

func runTasksComplete(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	taskID := args[0]

	if err := api.NewTasksService(client).Complete(context.Background(), taskID); err != nil {
		return fmt.Errorf("failed to complete task: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "taskId": "%s", "message": "Task completed successfully"}`, taskID)
	} else {
		fmt.Printf("Task %s completed successfully\n", taskID)
	}
	return nil
}

func runTasksDelete(cmd *cobra.Command, args []string) error {
	client, _, err := newClient()
	if err != nil {
		return err
	}

	taskID := args[0]

	if err := api.NewTasksService(client).Delete(context.Background(), taskID); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	out := getOutputFormat()
	if out == "json" {
		fmt.Printf(`{"status": "success", "taskId": "%s", "message": "Task deleted successfully"}`, taskID)
	} else {
		fmt.Printf("Task %s deleted successfully\n", taskID)
	}
	return nil
}