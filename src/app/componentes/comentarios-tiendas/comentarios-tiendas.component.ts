import { Component, OnInit } from '@angular/core';
import {baseURL} from "../../API/api";
import {ActivatedRoute} from "@angular/router";
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-comentarios-tiendas',
  templateUrl: './comentarios-tiendas.component.html',
  styleUrls: ['./comentarios-tiendas.component.css']
})
export class ComentariosTiendasComponent implements OnInit {

  constructor(private activedRoute: ActivatedRoute, private http: HttpClient) { }
  conversion
  comentarios: any[]
  ccomentario:any={
    Cui:5768581994586,
    Comentario:null,
  }
  ngOnInit(): void {

    this.http.get(baseURL + '/mostrarcomentarios')
      .subscribe(data => {
        this.conversion = data
        this.comentarios = this.conversion;
      });
  }
  ucomentario(){
    this.http.post('http://localhost:3000/crearcomentarios', this.ccomentario).toPromise().then((data:any)=>{
      console.log(data);
      // this.json=JSON.stringify(data.json);
    });
    window.location.href = 'http://localhost:4200/comentarios-tiendas';
  }
}
