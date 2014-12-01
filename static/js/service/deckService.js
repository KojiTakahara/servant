var service = angular.module('deckService', []);
service.factory('deckService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.validate = function(deck) {
    var lrigDeck = deck.Lrig,
        mainDeck = deck.Main,
        lrigNum = 0,
        mainNum = 0,
        burstNum = 0,
        alerts = [];
    for (var i in lrigDeck) {
      lrigNum += lrigDeck[i].Num;
    }
    for (var i in mainDeck) {
      mainNum += mainDeck[i].Num;
      burstNum += mainDeck[i].Bursted ? mainDeck[i].Num : 0;
    }
    if (isEmpty(deck.Title)) {
      alerts.push({ type: 'warning', msg: 'デッキ名を入力してください。' });
    }
    if (10 < lrigNum) {
      alerts.push({ type: 'warning', msg: 'ルリグデッキは合計10枚までにしてください。' });
    }
    if (mainNum !== 40) {
      alerts.push({ type: 'warning', msg: 'メインデッキは合計40枚にしてください。' });
    }
    if (burstNum !== 20) {
      alerts.push({ type: 'warning', msg: 'ライフバーストを持つカードは合計20枚にしてください。' });
    }
    return alerts;
  };

  var isEmpty = function(obj) {
    return obj === undefined || obj === "";
  };

  return service;

}]);