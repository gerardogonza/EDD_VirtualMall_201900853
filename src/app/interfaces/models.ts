export interface Pedido{
    productos: ProducotoPedido[];
    precioTotal:string;
    fecha: Date;
}
export interface ProducotoPedido{
    producto: Producto;
    cantidad: number;
}
export interface Producto{
    Nombre:string;
    Codigo: number;
    Descripcion: string;
    Precio: number;
    Cantidad: number;
}