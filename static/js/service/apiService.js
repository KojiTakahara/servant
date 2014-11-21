var service = angular.module('apiService', []);
service.factory('cardService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.getIllustrator = function() {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/illustrator',
      cache: true
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
      url: '/api/constraint',
      cache: true
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
      url: '/api/product',
      cache: true
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
      url: '/api/type',
      cache: true
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
      params: params,
      cache: true
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
      url: '/api/card/' + expansion,
      cache: true
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
      url: '/api/card/' + expansion + '/' + no,
      cache: true
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
    params.Lrig = [];
    params.Main = [];
    for (var i = 0; i < 10; i++) {
      if (!angular.isUndefined(params.lrig[i])) {
        for (var j = 0; j < params.lrig[i].num; j++) {
          params.Lrig.push(params.lrig[i]);
        }
      } else {
        params.Lrig.push({ KeyName: "" });
      }
    }
    for (var i in params.main) {
      for (var j = 0; j < params.main[i].num; j++) {
        params.Main.push(params.main[i]);
      }
    }
    if (params.Id) {
      // 更新
    } else {
      $http.post('/api/deck', params).success(function(data, status, headers, config) {
        deferred.resolve(data);
      }).error(function(data, status, headers, config) {
        deferred.reject(data);
      });
    }

    return deferred.promise;
  };

  return service;
}]);