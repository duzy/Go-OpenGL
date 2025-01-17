package glu

// #cgo darwin LDFLAGS: -framework OpenGL
// #cgo windows LDFLAGS: -lglu32
// #cgo linux LDFLAGS: -lGLU
//
// #ifdef __APPLE__
// # include <OpenGL/glu.h>
// #else
// # include <GL/glu.h>
// #endif
import "C"
import "gl"

type GLUquadric C.GLUquadric

func Build2DMipmaps(target gl.GLenum, internalFormat int, width, height int, format gl.GLenum, data interface{}) int {
	t, p := gl.GetGLenumType(data)
	return int(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(t),
		p,
	))
}

func Perspective(fovy, aspect, zNear, zFar float64) {
	C.gluPerspective(
		C.GLdouble(fovy),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) {
	C.gluLookAt(
		C.GLdouble(eyeX),
		C.GLdouble(eyeY),
		C.GLdouble(eyeZ),
		C.GLdouble(centerX),
		C.GLdouble(centerY),
		C.GLdouble(centerZ),
		C.GLdouble(upX),
		C.GLdouble(upY),
		C.GLdouble(upZ),
	)
}

// Project object coordinates to screen coordinates.
func Project(objx, objy, objz float64, model, proj []float64, view []int32) (x, y, z float64) {
	var wx, wy, wz C.GLdouble
	C.gluProject(
		C.GLdouble(objx), C.GLdouble(objy), C.GLdouble(objz), (*C.GLdouble)(&model[0]),
		(*C.GLdouble)(&proj[0]), (*C.GLint)(&view[0]), &wx, &wy, &wz,
	)
	return float64(wx), float64(wy), float64(wz)
}

// Project screen coordinates to object coordinates.
func Unproject(wx, wy, wz float64, model, proj []float64, view []int32) (objx, objy, objz float64) {
	var ox, oy, oz C.GLdouble
	C.gluUnProject(
		C.GLdouble(wx), C.GLdouble(wy), C.GLdouble(wz), (*C.GLdouble)(&model[0]),
		(*C.GLdouble)(&proj[0]), (*C.GLint)(&view[0]), &ox, &oy, &oz,
	)
	return float64(ox), float64(oy), float64(oz)
}

// Project screen coordinates to object coordinates.
func Unproject4(wx, wy, wz, clipw float64, model, proj []float64, view []int32, near, far float64) (objx, objy, objz, objw float64) {
	var ox, oy, oz, ow C.GLdouble
	C.gluUnProject4(
		C.GLdouble(wx), C.GLdouble(wy), C.GLdouble(wz), C.GLdouble(clipw),
		(*C.GLdouble)(&model[0]), (*C.GLdouble)(&proj[0]), (*C.GLint)(&view[0]),
		C.GLdouble(near), C.GLdouble(far), &ox, &oy, &oz, &ow,
	)
	return float64(ox), float64(oy), float64(oz), float64(ow)
}

func PickMatrix(x, y, delx, dely float64, viewport []int32) {
	C.gluPickMatrix(C.GLdouble(x), C.GLdouble(y), C.GLdouble(delx),
		C.GLdouble(dely), (*C.GLint)(&viewport[0]))
}

func Ortho2D(left, right, bottom, top float64) {
	C.gluOrtho2D(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top))
}

func NewQuadric() (quad *GLUquadric) {
        quad = (*GLUquadric)(C.gluNewQuadric())
        return
}

func (quad *GLUquadric) Delete() {
        C.gluDeleteQuadric((*C.GLUquadric)(quad))
}

//func (quad *GLUquadric) QuadricCallback(which GLenum, CallBackFunc _GLUfuncptr) {
//}

func (quad *GLUquadric) QuadricDrawStyle(draw gl.GLenum) {
        C.gluQuadricDrawStyle((*C.GLUquadric)(quad), C.GLenum(draw))
}

func (quad *GLUquadric) QuadricNormals(normal gl.GLenum) {
        C.gluQuadricNormals((*C.GLUquadric)(quad), C.GLenum(normal))
}

func (quad *GLUquadric) QuadricOrientation(orientation gl.GLenum) {
        C.gluQuadricOrientation((*C.GLUquadric)(quad), C.GLenum(orientation))
}

func (quad *GLUquadric) QuadricTexture(texture gl.GLboolean) {
        C.gluQuadricTexture((*C.GLUquadric)(quad), C.GLboolean(texture))
}

func (quad *GLUquadric) Cylinder(base, top, height gl.GLdouble, slices, stacks gl.GLint) {
        C.gluCylinder((*C.GLUquadric)(quad), C.GLdouble(base), C.GLdouble(top), C.GLdouble(height), C.GLint(slices), C.GLint(stacks))
}

func (quad *GLUquadric) Disk(inner, outer gl.GLdouble, slices, loops gl.GLint) {
        C.gluDisk((*C.GLUquadric)(quad), C.GLdouble(inner), C.GLdouble(outer), C.GLint(slices), C.GLint(loops))
}

func (quad *GLUquadric) PartialDisk(inner, outer gl.GLdouble, slices, loops gl.GLint, start, sweep gl.GLdouble) {
        C.gluPartialDisk((*C.GLUquadric)(quad), C.GLdouble(inner), C.GLdouble(outer), C.GLint(slices), C.GLint(loops), C.GLdouble(start), C.GLdouble(sweep))
}

func (quad *GLUquadric) Sphere(radius gl.GLdouble, slices, stacks gl.GLint) {
        C.gluSphere((*C.GLUquadric)(quad), C.GLdouble(radius), C.GLint(slices), C.GLint(stacks))
}
