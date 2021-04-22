import { NullTemplateVisitor } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
@Component({
  selector: 'app-registrar',
  templateUrl: './registrar.component.html',
  styleUrls: ['./registrar.component.css']
})
export class RegistrarComponent implements OnInit {

  constructor(private http: HttpClient) { }
  Dpi1:string;
  registrarU:any={
    Dpi:null,
    Password:null,
    Correo:null,
    Nombre:null,
    Cuenta:"Usuario"
  }
  ngOnInit(): void {
  }
  uregistras(){
    this.http.post('http://localhost:3000/registrarusuario', this.registrarU).toPromise().then((data:any)=>{
      console.log(data);
      // this.json=JSON.stringify(data.json);
    });
    window.location.href = 'http://localhost:4200/inicio';
}
  }


