var service = angular.module('apiService', []);
service.factory('cardService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.getIllustrator = function() {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/illustrator'
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getConstraint = function() {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/constraint'
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getProduct = function() {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/product'
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getType = function() {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/type'
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.search = function(params) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/search',
      params: params
    }).success(function(data, status, headers, config) {
      for (var i in data) {
        var cost = [];
        setCost(cost, data[i].CostWhite, '白');
        setCost(cost, data[i].CostRed, '赤');
        setCost(cost, data[i].CostBlue, '青');
        setCost(cost, data[i].CostGreen, '緑');
        setCost(cost, data[i].CostBlack, '黒');
        setCost(cost, data[i].CostColorless, '無');
        data[i].Cost = cost;
      }
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    var setCost = function(cost, color, str) {
      for (var i = 0; i < color; i++) {
        cost.push(str);
      }
    };
    return deferred.promise;
  };

  service.getCardByExpansion = function(expansion) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/card/' + expansion
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getCardByExpansionAndNo = function(expansion, no) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/card/' + expansion + '/' + no
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  /** デッキ保存 **/
  service.saveDeck = function(params) {
    var deferred = $q.defer();
    $http.post('/api/deck', params).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  return service;
}]);