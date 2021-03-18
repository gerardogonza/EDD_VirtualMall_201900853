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
  rta:[];
  ngOnInit(): void {
    
  }
mostrarTiendas(){
  this.http.get('https://www.datos.gov.co/resource/xdk5-pm3f.json')
  .subscribe(data => { 
  this.conversion=data
  this.rta=this.conversion;
  console.log('prueba',this.rta)
  });
}



}
