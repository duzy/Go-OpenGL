#ifdef __APPLE__
# include <OpenGL/glu.h>
#else
# include <GL/glu.h>
#endif
enum
{
  /* QuadricDrawStyle */
  $POINT = GLU_POINT,
  $LINE = GLU_LINE,
  $FILL = GLU_FILL,
  $SILHOUETTE = GLU_SILHOUETTE,

  /* QuadricNormal */
  $SMOOTH = GLU_SMOOTH,
  $FLAT = GLU_FLAT,
  $NONE = GLU_NONE,

  /* QuadricOrientation */
  $OUTSIDE = GLU_OUTSIDE,
  $INSIDE = GLU_INSIDE,
};
