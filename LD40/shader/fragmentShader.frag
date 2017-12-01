precision mediump float;

varying vec2 v_tex;

uniform mat4 u;
uniform mat4 c;
uniform mat4 p;

uniform sampler2D texture;

void main() {
    vec4 col = texture2D(texture, v_tex);

    gl_FragColor = col;
    //gl_FragColor = vec4(gl_FragCoord.z, v_tex.x, v_tex.y, 1.0);
}
