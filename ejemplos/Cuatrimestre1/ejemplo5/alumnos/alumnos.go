package alumnos

type Alumno struct {
	Nombre   string
	Legajo   int64
	Promedio float32
}

func (al *Alumno) Modificar() {
	al.Nombre = "Lucs"
}
