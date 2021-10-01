package memento

import "fmt"

// Editor is the Originator
type Editor struct {
	content string
}

func (e *Editor) CreateState() EditorState {
	return NewEditorState(e.content)
}

func (e *Editor) Restore(state EditorState) {
	e.content = state.GetContent()
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) SetContent(content string) {
	e.content = content
}

// EditorState is the Memento
type EditorState struct {
	content string
}

func NewEditorState(content string) EditorState {
	return EditorState{content: content}
}

func (e EditorState) GetContent() string {
	return e.content
}

// History is the Caretaker to decouple Editor and EditorState
// to satisfy Single Responsibility Principle.
type History struct {
	states []EditorState
}

func (h *History) Push(state EditorState) {
	h.states = append(h.states, state)
}

func (h *History) Pop() EditorState {
	lastIndex := len(h.states) - 1
	lastState := h.states[lastIndex]
	h.states = h.states[:lastIndex]
	return lastState
}

func Run() {
	editor := Editor{}
	history := History{states: []EditorState{}}

	editor.SetContent("a")
	history.Push(editor.CreateState())

	editor.SetContent("b")
	history.Push(editor.CreateState())

	editor.SetContent("c")
	editor.Restore(history.Pop())

	fmt.Println(editor.content)
}
