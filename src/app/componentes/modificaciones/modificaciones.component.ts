import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
@Component({
  selector: 'app-modificaciones',
  templateUrl: './modificaciones.component.html',
  styleUrls: ['./modificaciones.component.css']
})
export class ModificacionesComponent implements OnInit {

  constructor(private http: HttpClient) { }
  conversion
  tiendas:any=[];
  ngOnInit(): void {
  }

mostrarTiendas(){
  this.http.get('http://localhost:3000/mostrartiendas')
  .subscribe(data => {
    this.conversion=data
    this.tiendas=this.conversion;
    console.log('Tiendas',this.tiendas)
  });
}
}
