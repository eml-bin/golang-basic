# Go Modules

#### Commands

`go get`. Permite instalar paquetes
`go mod vendor`. Crear una carpeta <vendor>, es el código adicional que se usa en tu programa
`go mod tidy`. Hace una limpieza de aquellos paquetes que no utilices
`go mod edit -replace ...`. Puede reemplazar una paquetería por otra que se haya modificado por necesidad del proyecto
`go mod verify`. Verifica los modulos del proyecto actual