import syntax from './syntax';
import Prism from 'prismjs';

describe('Vali syntax', () => {
  it('should highlight Vali query correctly', () => {
    expect(Prism.highlight('{key="val#ue"}', syntax, 'vali')).toBe(
      '<span class="token context-labels"><span class="token punctuation">{</span><span class="token label-key attr-name">key</span>=<span class="token label-value attr-value">"val#ue"</span></span><span class="token punctuation">}</span>'
    );
    expect(Prism.highlight('{key="#value"}', syntax, 'vali')).toBe(
      '<span class="token context-labels"><span class="token punctuation">{</span><span class="token label-key attr-name">key</span>=<span class="token label-value attr-value">"#value"</span></span><span class="token punctuation">}</span>'
    );
    expect(Prism.highlight('{key="value#"}', syntax, 'vali')).toBe(
      '<span class="token context-labels"><span class="token punctuation">{</span><span class="token label-key attr-name">key</span>=<span class="token label-value attr-value">"value#"</span></span><span class="token punctuation">}</span>'
    );
    expect(Prism.highlight('#test{key="value"}', syntax, 'vali')).toBe(
      '<span class="token comment">#test{key="value"}</span>'
    );
    expect(Prism.highlight('{key="value"}#test', syntax, 'vali')).toBe(
      '<span class="token context-labels"><span class="token punctuation">{</span><span class="token label-key attr-name">key</span>=<span class="token label-value attr-value">"value"</span></span><span class="token punctuation">}</span><span class="token comment">#test</span>'
    );
    expect(Prism.highlight('{key="value"', syntax, 'vali')).toBe(
      '<span class="token context-labels"><span class="token punctuation">{</span><span class="token label-key attr-name">key</span>=<span class="token label-value attr-value">"value"</span></span>'
    );
  });
  it('should highlight functions in Vali query correctly', () => {
    expect(Prism.highlight('rate({key="value"}[5m])', syntax, 'vali')).toContain(
      '<span class="token function">rate</span>'
    );
    expect(Prism.highlight('avg_over_time({key="value"}[5m])', syntax, 'vali')).toContain(
      '<span class="token function">avg_over_time</span>'
    );
  });
  it('should highlight operators in Vali query correctly', () => {
    expect(Prism.highlight('{key="value"} |= "test"', syntax, 'vali')).toContain(
      '<span class="token operator"> |= </span>'
    );
    expect(Prism.highlight('{key="value"} |~"test"', syntax, 'vali')).toContain(
      '<span class="token operator"> |~</span>'
    );
  });
  it('should highlight pipe operations in Vali query correctly', () => {
    expect(Prism.highlight('{key="value"} |= "test" | logfmt', syntax, 'vali')).toContain(
      '<span class="token pipe-operator operator">|</span> <span class="token pipe-operations keyword">logfmt</span></span>'
    );
    expect(Prism.highlight('{key="value"} |= "test" | label_format', syntax, 'vali')).toContain(
      '<span class="token context-pipe"> <span class="token pipe-operator operator">|</span> <span class="token pipe-operations keyword">label_format</span></span>'
    );
  });
});
