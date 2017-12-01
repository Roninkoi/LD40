precision mediump float;

attribute vec4 a_pos;

attribute vec4 a_tex;

uniform mat4 u;
uniform mat4 c;
uniform mat4 p;

varying vec2 v_tex;

void main() {
    gl_Position = p*c*u*vec4(a_pos.xyz, 1.0);

    v_tex = vec2(a_tex.x*a_tex.z, a_tex.y*a_tex.w);
}
