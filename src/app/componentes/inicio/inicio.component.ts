import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-inicio',
  templateUrl: './inicio.component.html',
  styleUrls: ['./inicio.component.css']
})
export class InicioComponent implements OnInit {

  constructor(private http:HttpClient) { }
  conversion
  usuarios:any=[];
  ngOnInit(): void {
    this.http.get('http://localhost:3000/mostrarusuarios')
    .subscribe(data => {
      this.conversion=data
      this.usuarios=this.conversion;
      console.log('Usuarios',this.usuarios)
    });
  }
  Dpi:string; //Admin
  Password:string;
  DpiUser:string; //User
  PasswordUser:string;
  url = `http://localhost:3000/`;

  alogin(){

    if(this.Dpi=="1234567890101"&& this.Password=="1234"){
      console.log("Bienvenido Admin");
      window.location.href = 'http://localhost:4200/cargarjson';
    }
  for(var i = 0; i < this.conversion.length; i++){
    if(this.conversion[i].Dpi==this.Dpi&&this.conversion[i].Password==this.Password&&this.conversion[i].Cuenta=="Admin"){

      window.location.href = 'http://localhost:4200/cargarjson';
    }
  }

  alert('Datos Erroneos');
  return 0
}
ulogin(){

  for(var i = 0; i < this.conversion.length; i++){
    if(this.conversion[i].Dpi==this.DpiUser&&this.conversion[i].PasswordUser==this.Password&&this.conversion[i].Cuenta=="Usuario"){
      window.location.href = 'http://localhost:4200/modificaciones';
    }
  }
  alert('Datos Erroneos');
  return 0
}
}
