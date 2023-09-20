/* eslint-disable no-magic-numbers */

import { CommonConstant, CommonError } from '../constant/constants';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

import { Injectable } from '@angular/core';
import { LoaderService } from '../services/loader.service';
import { NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { environment } from '../../environments/environment';
import { NzMessageService } from 'ng-zorro-antd/message';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  protected apiEndpoint: string = environment.apiEndpoint;

  constructor(
    protected http: HttpClient,
    public router: Router,
    private modal: NzModalService,
    public loaderService: LoaderService,
    public messageSvc: NzMessageService
  ) { }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  protected getQueryFilter(queryObject: any): string {
    let query = '';
    const keyArray = Object.keys(queryObject);
    keyArray.forEach((key) => {
      if (queryObject[key] === '' || queryObject[key] === null) {
        return;
      }
      if (queryObject[key] instanceof Array) {
        const ObjectArray = queryObject[key];
        for (const eachObj of ObjectArray) {
          query += key + '=' + encodeURI(eachObj);
          query += '&';
        }
      } else {
        query += key + '=' + encodeURI(queryObject[key]);
        query += '&';
      }
    });
    if (query !== '') {
      // eslint-disable-next-line no-magic-numbers
      query = query.slice(0, -1);
      query = '?' + query;
    }
    return query;
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  protected convertEmptyValueToNull(queryObject: any): object | null | string | boolean | number {
    if (queryObject === null || queryObject === '') { return null; }
    if (typeof queryObject === 'number' || typeof queryObject === 'boolean' || typeof queryObject === 'string') {
      return queryObject;
    }
    const keyArray = Object.keys(queryObject);
    keyArray.forEach((key) => {
      if (typeof queryObject[key] === 'object' && queryObject[key] !== null) {
        if (Array.isArray(queryObject[key])) { /* empty */ } else {
          queryObject[key] = this.convertEmptyValueToNull(queryObject[key]);
        }
        return;
      }
      if (queryObject[key] === '') {
        queryObject[key] = null;
      }
    });
    return queryObject;
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  protected apiConnecter(method: string, queryUri: string, body: string | number | boolean | object | null): Observable<any> {
    const access_token = localStorage.getItem('access_token');
    const headerOptions = access_token ? new HttpHeaders()
      .set('Content-Type', 'application/json')
      .set('Authorization', access_token)
      :
      new HttpHeaders()
        .set('Content-Type', 'application/json');
    const options = { headers: headerOptions };

    switch (method) {
      case 'GET':
        queryUri += this.getQueryFilter(body);
        return this.http.get(queryUri, options).pipe(
          map(data => data),
          catchError(this.handleError));

      case 'POST':
        body = this.convertEmptyValueToNull(body);
        return this.http.post(queryUri, body, options).pipe(
          map(data => data),
          catchError(this.handleError));

      case 'PUT':
        body = this.convertEmptyValueToNull(body);
        return this.http.put(queryUri, body, options).pipe(
          map(data => data),
          catchError(this.handleError));

      case 'PATCH':
        body = this.convertEmptyValueToNull(body);
        return this.http.patch(queryUri, body).pipe(
          map(data => data),
          catchError(this.handleError));

      case 'DELETE':
        queryUri += this.getQueryFilter(body);
        return this.http.delete(queryUri, options).pipe(
          map(data => data),
          catchError(this.handleError));

      default:
        break;
    }
    return this.http.post(queryUri, body, options).pipe(
      map((data) => data),
      catchError(this.handleError));
  }

  protected handleError(error: HttpErrorResponse): Observable<HttpErrorResponse> {
    return throwError(() => error);
  }

  public handleErrorType(error: HttpErrorResponse) {
    if (error.status === 401) {
      this.errorDialog(CommonError.tokenExpired, 401);
    }
  }

  public handleMessageError(error: HttpErrorResponse) {
    if (error.status === 500) {
      this.messageErrorDialog(CommonError.internalServerError);
    } else if (error.status === 404) {
      const customError = this.customError(error);
      this.messageErrorDialog(customError);
    } else if (error.status === 400) {
      const customError = this.customError(error);
      this.messageErrorDialog(customError)
    } else if (error.status === 502) {
      this.messageErrorDialog(CommonError.badGateWay)
    }
  }

  public errorDialog(message: string, status: number) {
    this.modal.error({
      nzTitle: CommonError.title,
      nzOkText: CommonConstant.closeButton,
      nzContent: message,
      nzCentered: true,
      nzNoAnimation: true,
      nzMaskClosable: false,
      nzClosable: false,
      nzOnOk: () => status === 401 ? this.router.navigate(['signin']) : ''
    });
  }

  public messageErrorDialog(message: string) {
    this.messageSvc.error(
      message,
    )
  }

  public customError(error: HttpErrorResponse) {
    const errorMessage = error.error.error;
    const startIndex = errorMessage.lastIndexOf('=') + 2;
    const endIndex = errorMessage.length;
    const extractedMessage = errorMessage.substring(startIndex, endIndex);
    return extractedMessage
  }
}
