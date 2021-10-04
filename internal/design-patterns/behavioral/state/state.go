package state

// State Pattern allows object to alter its behavior when its internal state changes

// Canvas is the Context that satisfy Open Closed Principle
// Open for extension, cloused for modification
type Canvas struct {
	currentTool Tool
}

func (c *Canvas) MouseDown() {
	c.currentTool.MouseDown()
}

func (c *Canvas) MouseUp() {
	c.currentTool.MouseUp()
}

func (c *Canvas) GetCurrentTool() Tool {
	return c.currentTool
}

func (c *Canvas) SetCurrentTool(tool Tool) {
	c.currentTool = tool
}

// Tool is the State
type Tool interface {
	MouseDown()
	MouseUp()
}

// SelectionTool is the Concrete State
type SelectionTool struct{}

func (s *SelectionTool) MouseDown() {}
func (s *SelectionTool) MouseUp()   {}

// BrushTool is the Concrete State
type BrushTool struct{}

func (b *BrushTool) MouseDown() {}
func (b *BrushTool) MouseUp()   {}

func Run() {
	canvas := Canvas{}
	canvas.SetCurrentTool(&SelectionTool{})
	canvas.MouseDown()
	canvas.MouseUp()
}
