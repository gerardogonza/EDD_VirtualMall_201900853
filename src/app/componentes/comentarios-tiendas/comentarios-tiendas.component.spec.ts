import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ComentariosTiendasComponent } from './comentarios-tiendas.component';

describe('ComentariosTiendasComponent', () => {
  let component: ComentariosTiendasComponent;
  let fixture: ComponentFixture<ComentariosTiendasComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ComentariosTiendasComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ComentariosTiendasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
