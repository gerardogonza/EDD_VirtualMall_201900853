import { ThisReceiver } from '@angular/compiler';
import { Injectable } from '@angular/core';
import {Producto,Pedido,ProducotoPedido } from "../interfaces/models";
import { HttpClient } from "@angular/common/http";
@Injectable({
  providedIn: 'root'
})
export class CarritoService {
 private pedido:Pedido;
  constructor(private http: HttpClient) {
    this.loadCarrito();
   }
   loadCarrito(){
    this.pedido={
     productos:[],
     precioTotal:"12",
     fecha: new Date()
    }
  }
  addProducto(producto:Producto){
    const item=this.pedido.productos.find(productoPedido => {
      return(productoPedido.producto.Codigo===producto.Codigo)
    }); 
  if (item!==undefined) {
    item.cantidad ++;
  }
  else{
    const add: ProducotoPedido={
      cantidad: 1,
      producto,
    };
  this.pedido.productos.push(add);
  }
  console.log('en add pedido =>',this.pedido);
  }

  
  getCarrito(){
    
    return this.pedido
   
  }
  removeProducto(producto:Producto){

  }
  json;
  realizarPedido(){
    console.log('en add pedido =>',this.pedido.productos);
    this.http.post('http://localhost:3000/carrito', this.pedido).toPromise().then((data:any)=>{
      console.log(data);
      // this.json=JSON.stringify(data.json);
    });
   
  }
  
}

