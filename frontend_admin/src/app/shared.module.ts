import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { NgZorroAntdModule } from './ng-zorro-antd/ng-zorro-antd.module';
import { MenuComponent } from './components/menu/menu.component';
import { DepartmentCreateComponent } from './components/department-create/department-create.component';
import { CategoryCreateComponent } from './components/category-create/category-create.component';

@NgModule({
  declarations: [
    MenuComponent,
    DepartmentCreateComponent,
    CategoryCreateComponent
  ],
  imports: [
    NgZorroAntdModule,
    RouterModule,
    ReactiveFormsModule
  ],
  exports: [
    NgZorroAntdModule,
    MenuComponent,
    DepartmentCreateComponent,
    CategoryCreateComponent
  ]
})
export class SharedModule { }
