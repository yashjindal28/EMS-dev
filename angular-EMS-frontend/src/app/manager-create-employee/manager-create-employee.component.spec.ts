import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagerCreateEmployeeComponent } from './manager-create-employee.component';

describe('ManagerCreateEmployeeComponent', () => {
  let component: ManagerCreateEmployeeComponent;
  let fixture: ComponentFixture<ManagerCreateEmployeeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagerCreateEmployeeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ManagerCreateEmployeeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
