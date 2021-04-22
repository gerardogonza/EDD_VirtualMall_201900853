import { Component, Input, OnInit } from '@angular/core';
import { Pedido, ProducotoPedido } from 'src/app/interfaces/models';
import { CarritoService } from "../../services/carrito.service";
import { HttpClient } from "@angular/common/http";
@Component({
  selector: 'app-carrito-compras',
  templateUrl: './carrito-compras.component.html',
  styleUrls: ['./carrito-compras.component.css']
})
export class CarritoComprasComponent implements OnInit {
pedido:Pedido;

  constructor(private http: HttpClient) {

   }
   conversion
   carrito:any=[];

  ngOnInit(): void {
    this.http.get('http://localhost:3000/mostrarcarrito')
    .subscribe(data => {
      this.conversion=data
      this.carrito=this.conversion;
  });
}

generarPedido(){
  this.http.post('http://localhost:3000/rutamin', this.conversion).toPromise().then((data:any)=>{
      console.log(data);
      // this.json=JSON.stringify(data.json);
    });
    window.location.href = 'http://localhost:4200/modificaciones';
}
}
