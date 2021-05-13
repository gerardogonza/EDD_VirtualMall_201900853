import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ModificacionesComponent} from "./componentes/modificaciones/modificaciones.component";
import{ InicioComponent } from "./componentes/inicio/inicio.component";
import{InventarioComponent}from "./componentes/inventario/inventario.component";
import { CarritoComprasComponent } from "./componentes/carrito-compras/carrito-compras.component";
import { PedidoComponent } from './componentes/pedido/pedido.component';
import { CargaJSONComponent } from './componentes/carga-json/carga-json.component';
import { ReportesComponent } from './componentes/reportes/reportes.component';
import { RegistrarComponent } from './componentes/registrar/registrar.component';
import {ReportesmerkleComponent} from "./componentes/reportesmerkle/reportesmerkle.component";
import {ComentariosTiendasComponent} from "./componentes/comentarios-tiendas/comentarios-tiendas.component";
import {ComentariosProductosComponent} from "./componentes/comentarios-productos/comentarios-productos.component";
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
    path:'reportes',
    component:ReportesComponent
  },
  {
    path:'reportesmerkle',
    component:ReportesmerkleComponent
  },
  {
    path:'pedido/:nombre',
    component:PedidoComponent
  },
  {
    path:'registrar',
    component:RegistrarComponent
  },
  {
    path:'comentarios-tiendas',
    component:ComentariosTiendasComponent
  },
  {
    path:'comentarios-productos',
    component:ComentariosProductosComponent
  }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
