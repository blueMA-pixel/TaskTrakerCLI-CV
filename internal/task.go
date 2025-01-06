package internal

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "to-do"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"

	SORT_BY_STATUS        = 2
	SORT_BY_CREATION_TIME = 3
	SORT_BY_UPDATE_TIME   = 4

	DATE_FORMAT = "2006-01-02 15:04:05"
)

var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12"))

	TaskToDoStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("14"))

	TaskInProgressStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("11"))

	TaskDoneStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("2")).
			Strikethrough(true)
)

type Task struct {
	Id           int        `json:"ID"`
	Description  string     `json:"Description"`
	Status       TaskStatus `json:"Status"`
	CreationTime time.Time  `json:"CreationTime"`
	UpdateTime   time.Time  `json:"UpdateTime"`
}

func (t Task) String() string {
	return fmt.Sprintf(
		"ID: %-10d\nDescription: %-30s\nStatus: %-15s\nCreated: %-20s\nUpdated: %-20s\n",
		t.Id,
		t.Description,
		t.Status,
		t.CreationTime.Format(DATE_FORMAT),
		t.UpdateTime.Format(DATE_FORMAT),
	)
}

func NewTask(id int, description string) *Task {
	return &Task{
		Id:           id,
		Description:  description,
		Status:       TASK_STATUS_TODO,
		CreationTime: time.Now(),
		UpdateTime:   time.Now(),
	}
}

func ListTasksCommand(taskStatus *TaskStatus, sortBy int) error {
	tasks, err := readTasksFromDisk()

	if err != nil {
		return err
	}

	var rows [][]string

	for _, task := range tasks {
		if taskStatus == nil || task.Status != *taskStatus {
			rows = append(rows, []string{
				strconv.Itoa(task.Id),
				task.Description,
				string(task.Status),
				task.CreationTime.Format(DATE_FORMAT),
				task.UpdateTime.Format(DATE_FORMAT),
			})
		}
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i][sortBy] < rows[j][sortBy]
	})

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("56"))).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == -1:
				return HeaderStyle
			case rows[row][2] == string(TASK_STATUS_IN_PROGRESS):
				return TaskInProgressStyle
			case rows[row][2] == string(TASK_STATUS_TODO):
				return TaskToDoStyle
			case rows[row][2] == string(TASK_STATUS_DONE):
				return TaskDoneStyle
			default:
				return lipgloss.Style{}
			}
		}).
		Headers("ID", "Description", "Status", "Creation time", "Last update time")

	fmt.Println(t)
	return err
}

func AddTaskCommand(description string) error {
	tasks, err := readTasksFromDisk()

	if err != nil {
		return err
	}

	tasks = append(tasks, *NewTask(tasks[len(tasks)-1].Id+1, description))

	return writeTaskToDisk(&tasks)
}

func DeleteTaskCommand(id int) error {
	tasks, err := readTasksFromDisk()

	if err != nil {
		return err
	}

	var index int = -1
	for i, tasks := range tasks {
		if tasks.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("invalid task id")
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	return writeTaskToDisk(&tasks)
}

func UpdateTaskCommand(id int, description string) error {
	tasks, err := readTasksFromDisk()

	if err != nil {
		return err
	}

	var found bool
	for i, task := range tasks {
		if task.Id == id {
			found = true
			tasks[i].Description = description
			break
		}
	}

	if !found {
		return fmt.Errorf("invalid task id")
	}

	return writeTaskToDisk(&tasks)
}

func MarkCommand(id int, newStatus TaskStatus) error {
	tasks, err := readTasksFromDisk()

	if err != nil {
		return err
	}

	var found bool
	for i, task := range tasks {
		if task.Id == id {
			found = true
			tasks[i].Status = newStatus
			break
		}
	}

	if !found {
		return fmt.Errorf("inavalid task id")
	}

	return writeTaskToDisk(&tasks)
}
