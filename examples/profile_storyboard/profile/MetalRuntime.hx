package profile;

import domain.StoryCard;
import haxe.ds.List;

class MetalRuntime implements StoryboardRuntime {
  public function new() {}

  public function profileId():String {
    return "metal";
  }

  public function decorateTitle(title:String):String {
    return "[" + title + "]";
  }

  public function highlightTag(tag:String):String {
    return "metal-" + tag;
  }

  public function extraSignal(cards:List<StoryCard>):String {
    var highValue = 0;
    var count = cards.length;
    var i = 0;
    while (i < count) {
      var value = cards.pop();
      if (value == null) {
        break;
      }
      var card:StoryCard = cast value;
      if (card.points >= 5) {
        highValue++;
      }
      cards.add(card);
      i++;
    }
    return "interop_lane=typed+strict,high_value=" + highValue;
  }

  public function supportsVelocityHint():Bool {
    return true;
  }
}
