import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { en_US } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import en from '@angular/common/locales/en';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgZorroAntdModule } from './ng-zorro-antd/ng-zorro-antd.module';
import { LoadingComponent } from './components/loading/loading.component';
import { SharedModule } from './shared.module';
import { IconsProviderModule } from './icons-provider.module';
import { TeamCreateComponent } from './components/team-create/team-create.component';
import { CsvTableComponent } from './components/csv-table/csv-table.component';

registerLocaleData(en);

@NgModule({
  declarations: [
    AppComponent,
    LoadingComponent,
    TeamCreateComponent,
    CsvTableComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    NgZorroAntdModule,
    SharedModule,
    IconsProviderModule,
    ReactiveFormsModule
  ],
  providers: [
    { provide: NZ_I18N, useValue: en_US }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
