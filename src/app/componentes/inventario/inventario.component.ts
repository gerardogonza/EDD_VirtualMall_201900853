import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { HttpClient } from "@angular/common/http";
import { CarritoService } from "../../services/carrito.service";
import { Producto } from 'src/app/interfaces/models';
import { dashCaseToCamelCase, partitionArray } from '@angular/compiler/src/util';
@Component({
  selector: 'app-inventario',
  templateUrl: './inventario.component.html',
  styleUrls: ['./inventario.component.css']
})
export class InventarioComponent implements OnInit {

  constructor(private activedRoute: ActivatedRoute,private http: HttpClient, public carritoServices:CarritoService) { }
  conversion
  tiendas:Producto[]=[];

  ngOnInit(): void {
    const params =this.activedRoute.snapshot.params;
    if(params.nombre){
      this.http.get('http://localhost:3000/mostrarinventario/'+params.nombre)
      .subscribe(data => {
        this.conversion=data
        this.tiendas=this.conversion;
        // console.log('Tiendas',this.tiendas)
      });
    }
  }

addCarrito(index:number){
this.carritoServices.addProducto(this.conversion[index]);
}
getIndex(Codigo:number):number{
  for(var i = 0; i < this.conversion.length; i++){
    if(this.conversion[i].Codigo==Codigo){
      return i
    }
  }
  return 0 
}

}
