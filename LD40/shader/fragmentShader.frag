precision mediump float;

varying vec2 v_tex;

uniform mat4 u;
uniform mat4 c;
uniform mat4 p;

uniform sampler2D texture;

void main() {
    vec4 col = texture2D(texture, v_tex);

    col.r *= col.r*1.1;
    col.g *= col.g*1.05;
    col.b *= col.b;

    float depth = gl_FragCoord.z*1.505;

    depth *= depth*depth*depth;

    col.rgb *= min(1.0, max(0.0, 5.0 - depth));

    gl_FragColor = col;
}
