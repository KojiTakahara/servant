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

  /** 公開されているデッキを取得する **/
  service.getPublicDecks = function(limit, offset) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/deck',
      params: {
        limit: limit,
        offset: offset
      },
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getDeckById = function(id) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/deck/' + id,
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getDeckByUserId = function(userId) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/' + userId + '/deck',
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  }

  /** デッキ保存 **/
  service.saveDeck = function(params) {
    var deferred = $q.defer();
    params.UniqueLrigs = angular.copy(params.Lrig);
    params.UniqueMains = angular.copy(params.Main);
    params.OriginalLrigs = [];
    params.OriginalMains = [];
    for (var i = 0; i < 10; i++) {
      if (!angular.isUndefined(params.UniqueLrigs[i])) {
        for (var j = 0; j < params.UniqueLrigs[i].Num; j++) {
          params.OriginalLrigs.push(params.UniqueLrigs[i]);
        }
      } else {
        params.OriginalLrigs.push({ KeyName: "" });
      }
    }
    for (var i in params.UniqueMains) {
      for (var j = 0; j < params.UniqueMains[i].Num; j++) {
        params.OriginalMains.push(params.UniqueMains[i]);
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

  /** デッキ削除 **/
  service.deleteDeck = function(deck) {
    var deferred = $q.defer();
    $http({
      method: 'DELETE',
      url: '/api/deck/' + deck.Owner + '/' + deck.Id
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  return service;
}]);