import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { HttpClient } from "@angular/common/http";
import { baseURL } from '../../API/api'
@Component({
  selector: 'app-pedido',
  templateUrl: './pedido.component.html',
  styleUrls: ['./pedido.component.css']
})
export class PedidoComponent implements OnInit {

  constructor(private activedRoute: ActivatedRoute, private http: HttpClient) { }
  conversion
  pedido: any[]
  ngOnInit(): void {
  }
  mostraspedido() {
    const params = this.activedRoute.snapshot.params;
    if (params.nombre) {
      this.http.get(baseURL + '/mostrarpedido/' + params.nombre)
        .subscribe(data => {
          this.conversion = data
          this.pedido = this.conversion;
          console.log('pedido', this.pedido)
        });
    }
  }
}
