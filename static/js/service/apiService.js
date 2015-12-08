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
      cache: false
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
  service.getPublicDecks = function(params) {
    var deferred = $q.defer();
    params.white = (params.white === true) ? true : undefined;
    params.red = (params.red === true) ? true : undefined;
    params.blue = (params.blue === true) ? true : undefined;
    params.green = (params.green === true) ? true : undefined;
    params.black = (params.black === true) ? true : undefined;
    $http({
      method: 'GET',
      url: '/api/deck',
      params: params,
      cache: false
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  /**
   * IDに一致するデッキを取得
   * @param {long} id ID
   * @param {string} scope スコープ
   * @param {bool} mypage マイページ表示はtrue
   */
  service.getDeckById = function(id, scope, mypage) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/deck/' + id,
      params: {
        scope: scope,
        mypage: mypage
      },
      cache: false
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getDeckByUserId = function(userId, scope) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/' + userId + '/deck',
      params: {
        scope: scope,
      },
      cache: false
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
    $http.post('/api/deck', params).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
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

  service.getUser = function(userId) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/user/' + userId,
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.getTwitterUser = function(userId) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/twitter/user/' + userId,
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  return service;
}]);