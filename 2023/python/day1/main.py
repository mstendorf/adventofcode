with open("./data.txt") as f:
    lines = f.read().splitlines()
    numbers = []
    for line in lines:
        line_numbers = []
        for c in line:
            if c.isdigit():
                line_numbers.append(c)

        numbers.append(int(line_numbers[0] + line_numbers[-1]))

print(sum(numbers))
--- Startup times for process: Primary/TUI ---

times in msec
 clock   self+sourced   self:  sourced script
 clock   elapsed:              other lines

000.001  000.001: --- NVIM STARTING ---
000.104  000.104: event init
000.273  000.168: early init
000.439  000.166: locale set
000.488  000.049: init first window
012.215  011.727: inits 1
012.226  000.011: window checked
012.231  000.005: parsing arguments
013.013  000.031  000.031: require('vim.shared')
013.086  000.037  000.037: require('vim.inspect')
013.126  000.031  000.031: require('vim._options')
013.128  000.111  000.043: require('vim._editor')
013.129  000.203  000.062: require('vim._init_packages')
013.130  000.696: init lua interpreter
013.570  000.440: --- NVIM STARTED ---

--- Startup times for process: Embedded ---

times in msec
 clock   self+sourced   self:  sourced script
 clock   elapsed:              other lines

000.000  000.000: --- NVIM STARTING ---
000.080  000.079: event init
000.228  000.148: early init
000.499  000.271: locale set
000.534  000.035: init first window
000.809  000.275: inits 1
000.827  000.018: window checked
000.830  000.003: parsing arguments
001.202  000.026  000.026: require('vim.shared')
001.264  000.032  000.032: require('vim.inspect')
001.299  000.028  000.028: require('vim._options')
001.301  000.096  000.037: require('vim._editor')
001.302  000.146  000.024: require('vim._init_packages')
001.303  000.328: init lua interpreter
003.170  001.866: expanding arguments
003.179  000.009: inits 2
003.393  000.214: init highlight
003.394  000.001: waiting for UI
003.483  000.090: done waiting for UI
003.487  000.003: clear screen
003.547  000.009  000.009: require('vim.keymap')
004.067  000.578  000.569: require('vim._defaults')
004.069  000.004: init default mappings & autocommands
004.812  000.307  000.307: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/ftplugin.vim
005.131  000.228  000.228: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/indent.vim
006.360  000.554  000.554: require('mstendorf.set')
008.081  001.718  001.718: require('mstendorf.mappings')
009.135  000.650  000.650: require('lazy')
009.158  000.012  000.012: require('ffi')
009.201  000.024  000.024: require('vim.fs')
009.524  000.320  000.320: require('vim.uri')
009.531  000.371  000.027: require('vim.loader')
009.960  000.411  000.411: require('lazy.stats')
010.589  000.596  000.596: require('lazy.core.util')
010.793  000.202  000.202: require('lazy.core.config')
011.301  000.184  000.184: require('lazy.core.handler')
012.191  000.431  000.431: require('lazy.pkg')
012.200  000.616  000.185: require('lazy.core.meta')
012.205  000.902  000.286: require('lazy.core.plugin')
012.210  001.416  000.329: require('lazy.core.loader')
012.721  000.249  000.249: require('lazy.core.fragments')
017.675  000.123  000.123: require('lazy.core.handler.event')
017.678  000.255  000.132: require('lazy.core.handler.ft')
017.793  000.114  000.114: require('lazy.core.handler.keys')
017.930  000.134  000.134: require('lazy.core.handler.cmd')
018.593  000.210  000.210: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/filetype.lua
019.733  000.931  000.931: require('catppuccin')
020.472  000.166  000.166: require('catppuccin.lib.hashing')
022.663  001.502  001.502: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/catppuccin/colors/catppuccin.vim
023.056  000.133  000.133: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/plenary.nvim/plugin/plenary.vim
023.571  000.299  000.299: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/telescope.nvim/plugin/telescope.lua
024.554  000.974  000.974: require('telescope.builtin')
026.061  000.143  000.143: require('plenary.bit')
026.176  000.114  000.114: require('plenary.functional')
026.204  000.693  000.436: require('plenary.path')
026.213  000.902  000.209: require('plenary.strings')
026.349  000.135  000.135: require('telescope.deprecated')
026.854  000.003  000.003: require('vim.F')
026.867  000.175  000.171: require('plenary.log')
026.913  000.306  000.131: require('telescope.log')
027.398  000.079  000.079: require('plenary.compat')
027.404  000.282  000.202: require('plenary.job')
027.494  000.090  000.090: require('telescope.state')
027.518  000.604  000.232: require('telescope.utils')
027.521  001.171  000.261: require('telescope.sorters')
028.897  003.849  001.641: require('telescope.config')
029.258  000.131  000.131: require('plenary.window.border')
029.369  000.109  000.109: require('plenary.window')
029.461  000.091  000.091: require('plenary.popup.utils')
029.464  000.565  000.234: require('plenary.popup')
029.570  000.105  000.105: require('telescope.pickers.scroller')
029.679  000.108  000.108: require('telescope.actions.state')
029.820  000.140  000.140: require('telescope.actions.utils')
030.283  000.276  000.276: require('telescope.actions.mt')
030.378  000.557  000.281: require('telescope.actions.set')
030.608  000.115  000.115: require('telescope.config.resolve')
030.609  000.231  000.116: require('telescope.pickers.entry_display')
030.739  000.129  000.129: require('telescope.from_entry')
030.890  006.334  000.649: require('telescope.actions')
031.097  000.110  000.110: require('telescope._extensions')
031.099  000.208  000.098: require('telescope')
032.602  000.252  000.252: require('mason-core.path')
033.166  000.219  000.219: require('mason-core.functional')
033.387  000.087  000.087: require('mason-core.functional.data')
033.389  000.209  000.122: require('mason-core.functional.function')
033.483  000.087  000.087: require('mason-core.functional.relation')
033.571  000.084  000.084: require('mason-core.functional.logic')
033.576  000.973  000.375: require('mason-core.platform')
033.662  000.085  000.085: require('mason.settings')
033.664  001.679  000.368: require('mason')
034.144  000.217  000.217: require('mason-core.functional.list')
034.347  000.201  000.201: require('mason-core.functional.string')
034.359  000.581  000.163: require('mason.api.command')
034.541  000.176  000.176: require('mason-registry.sources')
036.175  000.534  000.534: require('cmp.utils.api')
036.754  000.096  000.096: require('cmp.types.cmp')
036.931  000.082  000.082: require('cmp.utils.misc')
036.936  000.181  000.099: require('cmp.types.lsp')
037.028  000.091  000.091: require('cmp.types.vim')
037.029  000.853  000.485: require('cmp.types')
037.108  000.078  000.078: require('cmp.utils.highlight')
037.270  000.090  000.090: require('cmp.utils.debug')
037.275  000.167  000.077: require('cmp.utils.autocmd')
037.424  001.962  000.330: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/nvim-cmp/plugin/cmp.lua
038.648  000.920  000.920: require('vim.filetype')
040.417  000.294  000.294: require('luasnip.util.types')
040.421  000.516  000.222: require('luasnip.util.ext_opts')
041.445  000.320  000.320: require('luasnip.util.lazy_table')
041.788  000.342  000.342: require('luasnip.extras.filetype_functions')
041.793  001.176  000.514: require('luasnip.default_config')
041.794  001.372  000.196: require('luasnip.session')
041.796  002.773  000.885: require('luasnip.config')
042.216  000.196  000.196: require('luasnip.util.util')
042.525  000.119  000.119: require('luasnip.nodes.key_indexer')
042.675  000.148  000.148: require('luasnip.util.feedkeys')
042.678  000.460  000.193: require('luasnip.nodes.util')
043.166  000.118  000.118: require('luasnip.session.snippet_collection.source')
043.326  000.159  000.159: require('luasnip.util.table')
043.459  000.133  000.133: require('luasnip.util.auto_table')
043.464  000.785  000.376: require('luasnip.session.snippet_collection')
044.454  000.243  000.243: require('luasnip.util.select')
044.838  000.382  000.382: require('luasnip.util.time')
045.151  001.524  000.899: require('luasnip.util._builtin_vars')
045.238  001.774  000.249: require('luasnip.util.environ')
045.451  000.212  000.212: require('luasnip.util.extend_decorator')
045.793  000.173  000.173: require('luasnip.util.path')
046.164  000.238  000.238: require('luasnip.util.log')
046.166  000.373  000.135: require('luasnip.loaders.util')
046.509  000.341  000.341: require('luasnip.loaders.data')
046.703  000.194  000.194: require('luasnip.loaders.fs_watchers')
046.705  001.253  000.172: require('luasnip.loaders')
046.715  004.916  000.236: require('luasnip')
046.724  009.154  000.546: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/LuaSnip/plugin/luasnip.lua
046.909  000.134  000.134: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/LuaSnip/plugin/luasnip.vim
047.959  000.157  000.157: require('vim.lsp.log')
048.629  000.668  000.668: require('vim.lsp.protocol')
049.091  000.165  000.165: require('vim.lsp._snippet_grammar')
049.188  000.096  000.096: require('vim.highlight')
049.196  000.565  000.304: require('vim.lsp.util')
049.379  000.084  000.084: require('vim.lsp.sync')
049.381  000.184  000.100: require('vim.lsp._changetracking')
049.644  000.205  000.205: require('vim.lsp.rpc')
049.705  002.138  000.360: require('vim.lsp')
049.932  000.226  000.226: require('lspconfig.util')
049.945  002.571  000.207: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/nvim-lspconfig/plugin/lspconfig.lua
050.254  000.244  000.244: require('vim.lsp.handlers')
050.891  000.136  000.136: require('cmp.utils.char')
050.895  000.248  000.112: require('cmp.utils.str')
051.589  000.192  000.192: require('cmp.utils.buffer')
051.593  000.561  000.369: require('cmp.utils.keymap')
051.594  000.698  000.136: require('cmp.utils.feedkeys')
052.393  000.492  000.492: require('cmp.config.mapping')
052.539  000.145  000.145: require('cmp.utils.cache')
053.197  000.506  000.506: require('cmp.config.compare')
053.198  000.658  000.152: require('cmp.config.default')
053.214  001.439  000.144: require('cmp.config')
053.220  001.626  000.187: require('cmp.utils.async')
053.569  000.202  000.202: require('cmp.utils.pattern')
053.572  000.350  000.148: require('cmp.context')
054.839  000.318  000.318: require('cmp.utils.snippet')
054.985  000.139  000.139: require('cmp.matcher')
054.993  000.711  000.254: require('cmp.entry')
054.998  001.426  000.715: require('cmp.source')
055.439  000.168  000.168: require('cmp.utils.event')
056.017  000.128  000.128: require('cmp.utils.options')
056.020  000.340  000.213: require('cmp.utils.window')
056.021  000.582  000.241: require('cmp.view.docs_view')
056.495  000.473  000.473: require('cmp.view.custom_entries_view')
056.705  000.209  000.209: require('cmp.view.wildmenu_entries_view')
056.821  000.115  000.115: require('cmp.view.native_entries_view')
056.941  000.119  000.119: require('cmp.view.ghost_text_view')
056.950  001.951  000.285: require('cmp.view')
057.299  006.846  000.548: require('cmp.core')
057.483  000.123  000.123: require('cmp.config.sources')
057.675  000.190  000.190: require('cmp.config.window')
057.704  007.449  000.289: require('cmp')
058.684  000.823  000.823: require('cmp_nvim_lsp.source')
058.685  000.980  000.157: require('cmp_nvim_lsp')
059.283  000.233  000.233: require('mason-core.log')
059.645  000.362  000.362: require('mason-lspconfig.settings')
059.648  000.821  000.227: require('mason-lspconfig')
059.883  000.124  000.124: require('mason-lspconfig.notify')
059.886  000.214  000.090: require('mason-lspconfig.lspconfig_hook')
060.601  000.118  000.118: require('mason-core.functional.table')
060.645  000.758  000.640: require('mason-lspconfig.mappings.server')
061.001  000.204  000.204: require('mason-core.EventEmitter')
061.494  000.492  000.492: require('mason-core.optional')
061.981  000.337  000.337: require('mason-core.async')
062.162  000.177  000.177: require('mason-core.async.uv')
062.171  000.675  000.161: require('mason-core.fs')
062.192  001.542  000.170: require('mason-registry')
062.397  000.203  000.203: require('mason-lspconfig.server_config_extensions')
062.919  000.291  000.291: require('lspconfig.async')
062.920  000.523  000.232: require('lspconfig.configs')
063.332  000.411  000.411: require('lspconfig.configs.omnisharp')
063.548  000.192  000.192: require('mason-lspconfig.ensure_installed')
064.203  000.103  000.103: require('mason-core.result')
064.610  000.165  000.165: require('mason-core.process')
064.778  000.167  000.167: require('mason-core.spawn')
064.781  000.446  000.114: require('mason-core.managers.powershell')
064.878  000.096  000.096: require('mason.version')
064.880  000.675  000.133: require('mason-core.fetch')
064.996  000.115  000.115: require('mason-core.providers')
065.461  000.141  000.141: require('mason-core.purl')
065.499  000.416  000.275: require('mason-core.package')
065.924  000.143  000.143: require('mason-core.installer.registry.expr')
065.929  000.324  000.181: require('mason-core.installer.registry.link')
067.089  000.129  000.129: require('mason-core.receipt')
067.101  000.283  000.153: require('mason-core.installer.context')
067.183  000.081  000.081: require('mason-core.async.control')
067.263  000.079  000.079: require('mason-core.installer.linker')
067.266  000.644  000.201: require('mason-core.installer')
067.273  001.193  000.549: require('mason-core.installer.managers.std')
067.274  001.345  000.152: require('mason-core.installer.registry.schemas')
067.355  000.080  000.080: require('mason-core.installer.registry.util')
067.360  001.860  000.111: require('mason-core.installer.registry')
067.361  002.364  000.088: require('mason-registry.sources.util')
067.366  003.806  000.549: require('mason-registry.sources.github')
070.992  000.091  000.091: require('mason-core.functional.number')
071.005  000.196  000.105: require('mason-lspconfig.api.command')
071.168  000.105  000.105: require('lspconfig')
071.357  000.188  000.188: require('lspconfig.configs.lua_ls')
072.885  000.135  000.135: require('lspconfig.manager')
072.985  000.096  000.096: require('lspconfig.configs.rust_analyzer')
073.452  000.093  000.093: require('lspconfig.configs.gopls')
073.823  000.079  000.079: require('lspconfig.configs.pyright')
074.793  000.320  000.320: require('vim.diagnostic')
075.611  000.172  000.172: require('rose-pine.config')
075.613  000.690  000.519: require('rose-pine')
076.981  000.161  000.161: require('treesitter-context.config')
076.989  000.924  000.763: require('treesitter-context')
078.204  000.090  000.090: require('vim.treesitter.language')
078.288  000.082  000.082: require('vim.func')
078.375  000.085  000.085: require('vim.func._memoize')
078.384  000.549  000.292: require('vim.treesitter.query')
078.474  000.088  000.088: require('vim.treesitter._range')
078.506  001.063  000.426: require('vim.treesitter.languagetree')
078.509  001.273  000.209: require('vim.treesitter')
078.706  000.196  000.196: require('vim.treesitter.highlighter')
078.788  000.081  000.081: require('treesitter-context.util')
078.791  001.774  000.225: require('treesitter-context.render')
078.814  002.909  000.211: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/nvim-treesitter-context/plugin/treesitter-context.lua
079.809  000.135  000.135: require('nvim-treesitter.compat')
081.551  001.496  001.496: require('nvim-treesitter.parsers')
081.684  000.132  000.132: require('nvim-treesitter.utils')
081.689  001.771  000.143: require('nvim-treesitter.ts_utils')
081.694  001.884  000.113: require('nvim-treesitter.tsrange')
081.792  000.098  000.098: require('nvim-treesitter.caching')
081.802  002.275  000.158: require('nvim-treesitter.query')
081.811  002.504  000.229: require('nvim-treesitter.configs')
081.814  002.758  000.254: require('nvim-treesitter-textobjects')
082.281  000.143  000.143: require('nvim-treesitter.info')
082.412  000.131  000.131: require('nvim-treesitter.shell_command_selectors')
082.437  000.524  000.250: require('nvim-treesitter.install')
082.534  000.096  000.096: require('nvim-treesitter.statusline')
082.633  000.099  000.099: require('nvim-treesitter.query_predicates')
082.635  000.820  000.101: require('nvim-treesitter')
082.845  000.121  000.121: require('nvim-treesitter.textobjects.shared')
082.849  000.210  000.088: require('nvim-treesitter.textobjects.select')
083.044  000.090  000.090: require('nvim-treesitter.textobjects.attach')
083.157  000.112  000.112: require('nvim-treesitter.textobjects.repeatable_move')
083.164  000.299  000.097: require('nvim-treesitter.textobjects.move')
083.421  000.086  000.086: require('nvim-treesitter.textobjects.swap')
083.566  000.097  000.097: require('nvim-treesitter.textobjects.lsp_interop')
083.578  004.641  000.370: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/nvim-treesitter-textobjects/plugin/nvim-treesitter-textobjects.vim
084.500  000.306  000.306: require('nvim-ts-autotag._log')
084.636  000.135  000.135: require('nvim-ts-autotag.config.init')
085.170  000.173  000.173: require('nvim-ts-autotag.utils')
085.176  000.380  000.207: require('nvim-ts-autotag.config.ft')
085.183  000.546  000.166: require('nvim-ts-autotag.config.plugin')
085.210  001.147  000.159: require('nvim-ts-autotag.internal')
085.212  001.489  000.343: require('nvim-ts-autotag')
085.213  001.538  000.049: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/nvim-ts-autotag/plugin/nvim-ts-autotag.lua
086.210  000.875  000.875: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/nvim-treesitter/plugin/nvim-treesitter.lua
087.424  000.181  000.181: require('tokyonight.config')
087.426  000.579  000.398: require('tokyonight')
087.688  000.146  000.146: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/editorconfig.lua
087.914  000.196  000.196: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/gzip.vim
088.080  000.137  000.137: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/man.lua
088.919  000.278  000.278: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/pack/dist/opt/matchit/plugin/matchit.vim
088.929  000.819  000.542: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/matchit.vim
089.178  000.204  000.204: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/matchparen.vim
089.529  000.281  000.281: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/netrwPlugin.vim
089.665  000.097  000.097: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/osc52.lua
089.827  000.126  000.126: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/rplugin.vim
090.020  000.156  000.156: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/shada.vim
090.205  000.149  000.149: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/spellfile.vim
090.356  000.114  000.114: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/tarPlugin.vim
090.593  000.200  000.200: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/tohtml.lua
090.798  000.168  000.168: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/tutor.vim
090.941  000.104  000.104: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/plugin/zipPlugin.vim
091.706  000.515  000.515: require('cmp_luasnip')
091.730  000.674  000.159: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/cmp_luasnip/after/plugin/cmp_luasnip.lua
092.399  000.466  000.466: require('cmp_cmdline')
092.407  000.593  000.127: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/cmp-cmdline/after/plugin/cmp_cmdline.lua
092.633  000.107  000.107: require('cmp_path')
092.643  000.173  000.066: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/cmp-path/after/plugin/cmp_path.lua
093.149  000.083  000.083: require('cmp_buffer.timer')
093.152  000.180  000.097: require('cmp_buffer.buffer')
093.153  000.277  000.097: require('cmp_buffer.source')
093.154  000.361  000.083: require('cmp_buffer')
093.191  000.447  000.087: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/cmp-buffer/after/plugin/cmp_buffer.lua
093.385  000.071  000.071: sourcing /Users/martinstendorf/.local/share/nvimnew/lazy/cmp-nvim-lsp/after/plugin/cmp_nvim_lsp.lua
093.436  085.353  019.484: require('mstendorf.lazy_init')
093.438  088.021  000.396: require('mstendorf')
093.439  088.180  000.159: sourcing /Users/martinstendorf/.config/nvimnew/init.lua
093.444  000.660: sourcing vimrc file(s)
093.759  000.025  000.025: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/filetype.lua
094.204  000.098  000.098: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/syntax/synload.vim
094.247  000.406  000.308: sourcing /opt/homebrew/Cellar/neovim/0.10.2_1/share/nvim/runtime/syntax/syntax.vim
094.254  000.379: inits 3
095.383  001.130: reading ShaDa
095.856  000.168  000.168: require('luasnip.util.directed_graph')
096.237  000.380  000.380: require('luasnip.session.enqueueable_operations')
096.241  000.780  000.232: require('luasnip.loaders.from_lua')
097.206  000.156  000.156: require('luasnip.util.events')
097.212  000.362  000.206: require('luasnip.nodes.node')
097.373  000.161  000.161: require('luasnip.nodes.insertNode')
097.537  000.162  000.162: require('luasnip.nodes.textNode')
097.935  000.398  000.398: require('luasnip.util.mark')
098.071  000.134  000.134: require('luasnip.util.pattern_tokenizer')
098.563  000.491  000.491: require('luasnip.util.dict')
099.295  000.637  000.637: require('luasnip.util.jsregexp')
099.297  000.733  000.096: require('luasnip.nodes.util.trig_engines')
099.314  002.710  000.270: require('luasnip.nodes.snippet')
099.316  002.796  000.086: require('luasnip.nodes.duplicate')
099.317  002.878  000.082: require('luasnip.loaders.snippet_cache')
099.321  003.079  000.200: require('luasnip.loaders.from_snipmate')
100.100  000.176  000.176: require('luasnip.util.parser.neovim_ast')
100.580  000.479  000.479: require('luasnip.util.str')
101.201  000.618  000.618: require('luasnip.util.jsregexp')
101.204  001.442  000.168: require('luasnip.util.parser.ast_utils')
101.313  000.108  000.108: require('luasnip.nodes.functionNode')
101.507  000.193  000.193: require('luasnip.nodes.choiceNode')
101.717  000.209  000.209: require('luasnip.nodes.dynamicNode')
101.808  000.089  000.089: require('luasnip.util.functions')
101.811  002.142  000.101: require('luasnip.util.parser.ast_parser')
101.992  000.180  000.180: require('luasnip.util.parser.neovim_parser')
102.004  002.418  000.096: require('luasnip.util.parser')
102.139  000.133  000.133: require('luasnip.nodes.snippetProxy')
102.323  000.183  000.183: require('luasnip.util.jsonc')
102.328  003.005  000.270: require('luasnip.loaders.from_vscode')
102.368  000.121: opening buffers
102.389  000.021: BufEnter autocommands
102.391  000.002: editing files in windows
102.457  000.066: VimEnter autocommands
102.566  000.110: UIEnter autocommands
102.567  000.001: before starting main loop
102.696  000.129: first screen update
102.697  000.001: --- NVIM STARTED ---

