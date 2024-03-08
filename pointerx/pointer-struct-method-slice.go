package pointerx

import "fmt"

type Subject struct {
	Name string
	SKS  uint
}

type Lecture struct {
	NIP     uint
	Nama    string
	Subject []Subject
}

func (lecture *Lecture) AssignSubject(subject *Subject) {
	lecture.Subject = append(lecture.Subject, *subject)
}

func (lecture *Lecture) PrintLecture() {

	transformSubject := []Subject{}

	for _, value := range lecture.Subject {
		transformSubject = append(transformSubject, Subject{
			Name: "Mata Kuliah: " + value.Name,
			SKS:  value.SKS,
		})
	}

	lecture.Subject = transformSubject

	fmt.Println(*lecture)

}

func Basic4() {
	lecture1 := &Lecture{NIP: 1556688222, Nama: "Fajar Pradana"}
	lecture1.AssignSubject(&Subject{"RPL", 5})
	lecture1.AssignSubject(&Subject{"Matdas", 4})
	lecture1.PrintLecture()
}
