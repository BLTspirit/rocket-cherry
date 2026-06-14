components {
  id: "bg_layer"
  component: "/Level/bg_layer.script"
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"bg_1_1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Assets/Sprites/bg.atlas\"\n"
  "}\n"
  ""
  position {
    x: 432.0
    y: 104.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bg_1_1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Assets/Sprites/bg.atlas\"\n"
  "}\n"
  ""
  position {
    x: 144.0
    y: 104.0
  }
}
