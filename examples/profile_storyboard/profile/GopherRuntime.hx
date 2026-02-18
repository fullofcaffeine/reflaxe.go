package profile;

import domain.StoryCard;
import haxe.ds.List;

class GopherRuntime implements StoryboardRuntime {
  public function new() {}

  public function profileId():String {
    return "gopher";
  }

  public function decorateTitle(title:String):String {
    return title + " [go]";
  }

  public function highlightTag(tag:String):String {
    return "go-" + tag;
  }

  public function extraSignal(cards:List<StoryCard>):String {
    var total = 0;
    var count = cards.length;
    var i = 0;
    while (i < count) {
      var value = cards.pop();
      if (value == null) {
        break;
      }
      var card:StoryCard = cast value;
      total += card.points;
      cards.add(card);
      i++;
    }
    return "interop_lane=typed,total_points=" + total;
  }

  public function supportsVelocityHint():Bool {
    return true;
  }
}
