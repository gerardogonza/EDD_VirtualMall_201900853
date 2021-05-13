
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ITiendas } from "../interfaces/tiendas";
@Injectable({
  providedIn: 'root',
})
export class TiendasService {
  constructor(private http: HttpClient) {}

  vertiendas() {
    return this.http.get('http://localhost:3000/cargartienda');
  }
API='http://localhost:3000/api';
  cargarTiendas(jsonEntrada:ITiendas):Observable<any>{
  return this.http.post(`${this.API}/cargatienda`,jsonEntrada);
  }


}
