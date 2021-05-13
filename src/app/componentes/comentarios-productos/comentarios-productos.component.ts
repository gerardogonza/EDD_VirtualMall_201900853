import { Component, OnInit,Input } from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {HttpClient} from "@angular/common/http";
import {baseURL} from "../../API/api";

@Component({
  selector: 'app-comentarios-productos',
  templateUrl: './comentarios-productos.component.html',
  styleUrls: ['./comentarios-productos.component.css']
})
export class ComentariosProductosComponent implements OnInit {

  constructor(private activedRoute: ActivatedRoute, private http: HttpClient) { }
  conversion
  comentarios: any[]
  ccomentario:any={
    Cui:5768581994586,
    Comentario:null,
  }
  ngOnInit(): void {

    this.http.get(baseURL + '/mostrarcomentariosproductos')
      .subscribe(data => {
        this.conversion = data
        this.comentarios = this.conversion;
      });
  }
  ucomentario(){
    this.http.post('http://localhost:3000/crearcomentariosproductos', this.ccomentario).toPromise().then((data:any)=>{
      console.log(data);
      // this.json=JSON.stringify(data.json);
    });
    window.location.href = 'http://localhost:4200/comentarios-productos';
  }
}
