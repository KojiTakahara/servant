var fil = angular.module('stringFilter', []);
fil.filter('color', function() {
  return function(input) {
    switch (input){
    case 'white':
      return '白';
      break;
    case 'red':
      return '赤';
      break;
    case 'blue':
      return '青';
      break;
    case 'green':
      return '緑';
      break;
    case 'black':
      return '黒';
      break;
    case 'colorless':
      return '無';
      break;
    default:
      return '';
      break;
    }
  };
});

fil.filter('productId', function() {
  return function(id) {
    var list = id.split(/([a-zA-Z]+)([0-9]+)/);
    if (list[1] === 'WD') {
      list[1] = 'WXD';
    }
    return list[1] + '-' + list[2];
  };
});