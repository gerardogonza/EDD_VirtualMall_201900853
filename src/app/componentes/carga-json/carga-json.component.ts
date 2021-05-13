import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-carga-json',
  templateUrl: './carga-json.component.html',
  styleUrls: ['./carga-json.component.css']
})
export class CargaJSONComponent  {

  cargatienda: String;
  cargarpedidido: String;
  cargarinventario: String;
  cargargrafo:String;
  cargarusuario:String;
  url = `http://localhost:3000/`;
  json;
  json1;
  json2;
  json3;
  json4;
  timer:any={
    tiempo:null
  }
  constructor(private http:HttpClient) {

   }

  // ngOnInit(): void {
  // }
  cargarTiendas(){
    this.http.post(this.url+`cargartienda`,this.cargatienda).toPromise().then((data:any)=>{
      console.log(data);
      this.json=JSON.stringify(data.json);
    });
}
cargarPedido(){
  this.http.post(this.url+`cargarpedido`,this.cargarpedidido).toPromise().then((data:any)=>{
    console.log(data);
    this.json1=JSON.stringify(data.json1);
  });
}
cargarInventario(){
  this.http.post(this.url+`cargarinventario`,this.cargarinventario).toPromise().then((data:any)=>{
    console.log(data);
    this.json2=JSON.stringify(data.json2);
  });
}
cargarGrafo(){
  this.http.post(this.url+`cargargrafo`,this.cargargrafo).toPromise().then((data:any)=>{
    console.log(data);
    this.json3=JSON.stringify(data.json3);
  });
}
cargarUsuarios(){
  this.http.post(this.url+`cargarusuarios`,this.cargarusuario).toPromise().then((data:any)=>{
    console.log(data);
    this.json4=JSON.stringify(data.json4);
  });
}
  configuracionTiempo(){
    this.http.post('http://localhost:3000/tiempo', this.timer).toPromise().then((data:any)=>{
      console.log(data);
      // this.json=JSON.stringify(data.json);
    });
    window.location.href = 'http://localhost:4200/cargarjson';
  }
  combrobarTiempo(){
    this.http.post('http://localhost:3000/comprobartiempo', this.timer).toPromise().then((data:any)=>{
      console.log(data);
      // this.json=JSON.stringify(data.json);
    });

  }
}
