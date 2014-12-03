var service = angular.module('deckService', []);
service.factory('deckService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.countLrig = function(deck) {
    var lrigNum = 0;
    for (var i in deck.Lrig) {
      lrigNum += deck.Lrig[i].Num;
    }
    return lrigNum;
  };

  service.countMain = function(deck) {
    var mainNum = 0;
    for (var i in deck.Main) {
      mainNum += deck.Main[i].Num;
    }
    return mainNum;
  };

  service.countBurst = function(deck) {
    var burstNum = 0;
    for (var i in deck.Main) {
      burstNum += deck.Main[i].Bursted ? deck.Main[i].Num : 0;
    }
    return burstNum;
  };

  service.validate = function(deck) {
    var alerts = [];
    if (isEmpty(deck.Title)) {
      alerts.push({ type: 'warning', msg: 'デッキ名を入力してください。' });
    }
    if (10 < service.countLrig(deck)) {
      alerts.push({ type: 'warning', msg: 'ルリグデッキは合計10枚までにしてください。' });
    }
    if (service.countMain(deck) !== 40) {
      alerts.push({ type: 'warning', msg: 'メインデッキは合計40枚にしてください。' });
    }
    if (service.countBurst(deck) !== 20) {
      alerts.push({ type: 'warning', msg: 'ライフバーストを持つカードは合計20枚にしてください。' });
    }
    return alerts;
  };

  var isEmpty = function(obj) {
    return obj === undefined || obj === "";
  };

  return service;

}]);