# Documento de cierre

## Que aprendimos

Aprendimos a organizar una API Go con separacion real entre handlers, services y repositories. Tambien practicamos autenticacion JWT, migraciones con GORM, pruebas unitarias con mocks y despliegue local con Docker Compose. El trabajo por modulos ayudo a entender como una entidad de un compañero afecta a otra: por ejemplo, una orden depende de usuarios y productos, y una mision de usuario depende de users y missions.

## Que hariamos distinto

Planificariamos antes los nombres de endpoints, structs y tablas para evitar inconsistencias entre español e ingles. Tambien definiriamos desde el inicio una politica de roles y permisos para no agregarla al final. Para una siguiente version convendria ampliar las pruebas de middleware y handlers de autenticacion.

## Proximos pasos del producto

El siguiente paso seria construir un frontend sencillo para estudiantes, agregar filtros de productos por categoria/precio, registrar imagenes de productos, mejorar la mensajeria entre compradores y vendedores, y crear un panel admin para categorias, insignias y misiones. Tambien se podria publicar la API en un hosting con PostgreSQL.
