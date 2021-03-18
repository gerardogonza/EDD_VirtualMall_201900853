import { Component, NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ModificacionesComponent} from "./componentes/modificaciones/modificaciones.component";
import{ InicioComponent } from "./componentes/inicio/inicio.component";
const routes: Routes = [
 
  {
    path:'inicio',
    component:InicioComponent
  },
  {
    path:'modificaciones',
    component:ModificacionesComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
