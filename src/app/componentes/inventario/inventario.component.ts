import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { HttpClient } from "@angular/common/http";
@Component({
  selector: 'app-inventario',
  templateUrl: './inventario.component.html',
  styleUrls: ['./inventario.component.css']
})
export class InventarioComponent implements OnInit {

  constructor(private activedRoute: ActivatedRoute,private http: HttpClient) { }
  conversion
  tiendas:any;
  ngOnInit(): void {
    const params =this.activedRoute.snapshot.params;
    if(params.nombre){
      this.http.get('http://localhost:3000/mostrarinventario/'+params.nombre)
      .subscribe(data => {
        this.conversion=data
        this.tiendas=this.conversion;
        console.log('Tiendas',this.tiendas)
      });
    }
  }

}
