import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ModificacionesComponent} from "./componentes/modificaciones/modificaciones.component";
import{ InicioComponent } from "./componentes/inicio/inicio.component";
import{InventarioComponent}from "./componentes/inventario/inventario.component";
import { CarritoComprasComponent } from "./componentes/carrito-compras/carrito-compras.component";
import { PedidoComponent } from './componentes/pedido/pedido.component';
import { CargaJSONComponent } from './componentes/carga-json/carga-json.component';
const routes: Routes = [
 
  {
    path:'inicio',
    component:InicioComponent
  },
  {
    path:'modificaciones',
    component:ModificacionesComponent
  },
  {
    path:'inventario/:nombre',
    component:InventarioComponent
  },
  {
    path:'carrito-compras',
    component:CarritoComprasComponent
  },
  {
    path:'cargarjson',
    component:CargaJSONComponent
  },
  {
    path:'pedido/:nombre',
    component:PedidoComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }