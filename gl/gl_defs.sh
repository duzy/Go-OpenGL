
echo "#include <GL/glew.h>
enum
{
" > gl_defs.c

sed -n 's/#define GL_\([0-9][a-zA-Z0-9_]*\) .*/\t$GL_\1 = GL_\1,/p' /usr/include/GL/glew.h | sort | uniq >> gl_defs.c
sed -n 's/#define GL_\([a-zA-Z][a-zA-Z_0-9]*\) .*/\t$\1 = GL_\1,/p' /usr/include/GL/glew.h | sort | uniq >> gl_defs.c

echo '};' >> gl_defs.c

godefs -g gl gl_defs.c > gl_defs.go
gofmt -w gl_defs.go
rm gl_defs.c

