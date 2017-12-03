precision mediump float;

varying vec2 v_tex;

uniform mat4 u;
uniform mat4 c;
uniform mat4 p;

uniform sampler2D texture;

void main() {
    vec4 col = texture2D(texture, v_tex);

    col += vec4(0.2, 0.14, 0.0, 0.0);

    float depth = gl_FragCoord.z;

    depth *= depth*depth*depth*depth*depth*depth; // lol

    col.rgb *= 1.0 - depth;

    gl_FragColor = col;
}
