import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http'; 
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ModificacionesComponent } from './componentes/modificaciones/modificaciones.component';
import { InventarioComponent } from './componentes/inventario/inventario.component';
import { InicioComponent } from './componentes/inicio/inicio.component';
import { CarritoComprasComponent } from './componentes/carrito-compras/carrito-compras.component';
import { PedidoComponent } from './componentes/pedido/pedido.component';

@NgModule({
  declarations: [
    AppComponent,
    ModificacionesComponent,
    InventarioComponent,
    InicioComponent,
    CarritoComprasComponent,
    PedidoComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
