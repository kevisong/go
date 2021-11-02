package visitor

import "fmt"

// The visitor pattern separates the algorithm from the objects
// When new object is added(in this case - new file types), all
// it should do is to implement the Accept(Visitor) method.
// When the Visit method needs to be changed, the original object
// will not be modified.

type Visitor interface {
	Visit(Visitable) error
}

// Compressor is the visitor
type Compressor struct{}

func (c *Compressor) VisitPDFFile() {
	fmt.Println("visiting pdf file")
}

func (c *Compressor) VisitPPTFile() {
	fmt.Println("visiting ppt file")
}

func (c *Compressor) Visit(v Visitable) error {
	switch t := v.(type) {
	case *PDFFile:
		c.VisitPDFFile()
	case *PPTFile:
		c.VisitPPTFile()
	default:
		return fmt.Errorf("type not supported: %s", t)
	}
	return nil
}

type Visitable interface {
	Accept(Visitor) error
}

// PDFFile is visitable
type PDFFile struct{}

func (p *PDFFile) Accept(visitor Visitor) error {
	return visitor.Visit(p)
}

// PDFFile is visitable
type PPTFile struct{}

func (p *PPTFile) Accept(visitor Visitor) error {
	return visitor.Visit(p)
}

func Run() {
	compressor := Compressor{}
	pdfFile := PDFFile{}
	compressor.Visit(&pdfFile)
	PPTFile := PPTFile{}
	compressor.Visit(&PPTFile)
}
