# Go Graphics Experiments

Writing an Entity Component System 3D engine with OpenGL.

Happy with progress so far, think the ECS is very very fast, optimises lookup speed over memory usage (lots of maps). Code is a bit dirty

Thinking of automating creation of helper functions like `entity.AddCameraComponent` through code gen from component models.

Not much graphics code, just a controllable cube and static camera. Working out best way to render. Grab renderables and store visible renderables in camera component, renderer system gets cameras and renders their renderables?
 
